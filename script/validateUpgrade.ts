import { exec } from 'child_process';
import * as process from 'process';
import yargs from 'yargs';
import * as fs from 'fs';
import { createWriteStream } from 'fs';
import * as goerliDeployData from './output/M1_deployment_goerli_2023_3_23.json'
import * as mainnetDeployData from './output/M1_deployment_mainnet_2023_6_9.json'
import "dotenv/config";

async function validateStorageSlots() {
  const args = await yargs
    .option('old', {
      alias: 'o',
      describe: 'Specify the filepath to the storage layout of the contract as it exists on chain',
      demandOption: true,
    })
    .option('new', {
      alias: 'n',
      describe: 'Specify the filepath to the storage layout of the contract to upgrade to',
      demandOption: true,
  }).argv;

  // Get csv files and format into mappings
  const oldLayoutCSV = fs.readFileSync(args.old as string, 'utf8');
  const oldLayout = formatCSV(oldLayoutCSV)
  const newLayoutCSV = fs.readFileSync(args.new as string, 'utf8');
  const newLayout = formatCSV(newLayoutCSV)

  // Assert that the new layout is not smaller than the old layout
  if (oldLayout.length > newLayout.length) {
    throw new Error('New storage layout is smaller than old storage layout');
  }

  // List of warnings
  const warnings = [];

  // List of errors
  const errors: Error[] = [];

  // Iterator to print out extra slots after search
  let index = 0;

  // Loop through new layout
  while (index < newLayout.length) {
    // Break if we have reached the end of the old layout
    if (index >= oldLayout.length) {
      break;
    }

    // Get item in slot
    const newEntry = newLayout[index];
    const oldEntry = oldLayout[index];

    if (newEntry.slot !== oldEntry.slot) {
      errors.push(new Error(`Slot ${newEntry.slot} has changed from ${oldEntry.slot}. Ensure that all old storage slots are present in new layout`));
    }

    if (newEntry.type !== oldEntry.type && oldEntry.name !== '__gap') { // Assume gaps are safe
      errors.push(new Error(`Slot ${newEntry.slot} has changed type from ${oldEntry.type} to ${newEntry.type}`));
    }

    if (newEntry.name !== oldEntry.name && oldEntry.name !== '__gap') {
      warnings.push(`Double check names: Slot ${newEntry.slot} has changed name from ${oldEntry.name} to ${newEntry.name}`);
    }

    // Mark slot as read
    oldLayout[index].read = true;

    // Increment index
    index++;
  }

  // Sanity check that all slots in old layout have been read
  for (let i = 0; i < oldLayout.length; i++) {
    if (!oldLayout[i].read) {
      errors.push(new Error(`Slot with name ${oldLayout[i].name} of type ${oldLayout[i].type} has been removed`));
    }
  }

  // Print new storage slots
  if (index < newLayout.length) {
    console.log('New storage slots:');
    for (let i = index; i < newLayout.length; i++) {
      console.log(`${newLayout[i].name} at slot ${newLayout[i].slot} of type ${newLayout[i].type}`);
    }
  }
  
  // Print warnings
  if (warnings.length > 0) {
    for (let i = 0; i < warnings.length; i++) {
      console.warn(warnings[i]);
    }
  }

  // Print errors
  if (errors.length > 0) {
    const aggregatedError = new Error("Errors found in storage layout");
    (aggregatedError as any).error = errors;
    throw aggregatedError;
  } else {
    console.log('Contract is upgrade safe');
  }
}

function formatCSV(csv: string) {
  const layoutFormatted = csv.split('\n')
  const storageSlots: StorageSlot[] = [];

  // Begin iterating when table begins, ignore any output before
  let storageBegin = 2; // storage begins on line 2 if there is no extra output
  for (let i = 0; i < layoutFormatted.length; i++) {
    if (layoutFormatted[i].includes('| Name')){
      break;
    }
    storageBegin++;
  }

  for (let i = storageBegin; i < layoutFormatted.length - 1; i++) {
    const data = layoutFormatted[i].split('|'); 
    
    // Extract the relevant data from each line
    const entry = {
      name: data[1].trim(),
      type: data[2].trim(),
      slot: Number(data[3].trim()),
      offset: Number(data[4].trim()),
      bytes: Number(data[5].trim()),
      value: Number(data[6].trim()),
      read: false
    }

    // Push the entry to the storageSlots array
    storageSlots.push(entry);
  }

  // Assert that the storageSlots array is sorted by slot
  for (let i = 0; i < storageSlots.length - 1; i++) {
    if (storageSlots[i].slot > storageSlots[i+1].slot) {
      console.error('Storage slots are not sorted by slot number');
      process.exit(1);
    }
  }

  return storageSlots;
}

validateStorageSlots();

type StorageSlot = {
  name: string,
  type: string,
  slot: number,
  offset: number,
  bytes: number,
  value: number,
  read: boolean
}