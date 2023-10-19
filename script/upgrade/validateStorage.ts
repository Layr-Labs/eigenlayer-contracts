import yargs, { string } from 'yargs';
import * as fs from 'fs';
import "dotenv/config";
import chalk from 'chalk';

// This function does basic validation of storage
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
    })
    .option('keep', {
      alias: 'k',
      describe: "Whether to keep the csv storage layout files"
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

  // Loop through new layout
  for(let slotNumber = 0; slotNumber < newLayout.length; slotNumber++) {
    // No more storage slots to compare with in old layout
    if (slotNumber >= oldLayout.length) {
      break;
    }

    // Mark slot as read
    oldLayout[slotNumber].read = true;

    // Get item in slot
    const newEntry = newLayout[slotNumber];
    const oldEntry = oldLayout[slotNumber];

    // Check that slot number is correct
    if (newEntry.slot !== oldEntry.slot) {
      throw new Error("Slot number mismatch");
    }

    // If old slot is empty and new slot is not, add an error - overwrote old storage with an empty slot
    if (newEntry.empty && !oldEntry.empty) {
      errors.push(new Error(`Slot ${newEntry.slot} has been incorrectly overridden`));
      continue;
    }

    // If either slot is empty, continue
    // Case1: newEntry.empty && oldEntry.empty -> fine
    // Case2: newEntry.empty && !oldEntry.empty -> error already added
    // Case3: !newEntry.empty && oldEntry.empty -> created new slot, will check with gaps
    if (newEntry.empty || oldEntry.empty) {
      continue;
    }

    // Remaining slots are now non-empty non-gaps
    // Check that slot name has not changed
    if (newEntry.name !== oldEntry.name) {
      warnings.push(`Double check names: Slot ${slotNumber} has changed name from ${oldEntry.name} to ${newEntry.name}`);
    }

    // Check that slot type has not changed
    if (newEntry.type !== oldEntry.type) {
      errors.push(new Error(`Slot ${slotNumber} has changed type from ${oldEntry.type} to ${newEntry.type}`));
    }
  }

  // Check gaps, starting from from the beginning of the old layout
  for(let slotNumber = 0; slotNumber < oldLayout.length; slotNumber++) {
    // Ignore non gap slots
    if(oldLayout[slotNumber].name !== '__gap') {
      continue;
    }

    // If the size is the same, continue
    if (oldLayout[slotNumber].type === newLayout[slotNumber].type) {
      continue;
    }

    // Gap is not present at the same slot in the new layout, find the next gap
    let newGapIndex = slotNumber;
    while (newLayout[newGapIndex].name !== '__gap' && newGapIndex < newLayout.length - 1) {
      newGapIndex++;
    }

    // Add error if no gap is found, we should have gaps at the end of all contracts
    if (newGapIndex === newLayout.length - 1 && newLayout[newGapIndex].name !== '__gap') {
      errors.push(new Error("No gap added to end of new storage layout"));
      continue;
    }

    // Get number of slots between gaps and extract the gap size
    const newSlots = newLayout[newGapIndex].slot - newLayout[slotNumber].slot;
    const oldGapSize = getGapSize(oldLayout[slotNumber].type);
    const newGapSize = getGapSize(newLayout[newGapIndex].type);

    // Check that gap has been properly resized
    if(newSlots + newGapSize !== oldGapSize) {
      // Okay to resize down if there is a non-empty slot after the gap and it's aligned to the 50th + 1 slot
      if(newSlots + newGapSize < oldGapSize && !newLayout[newGapIndex + newGapSize].empty && newLayout[newGapIndex + newGapSize].slot % 50 === 1){
        continue;
      } else{
        warnings.push(`Gap previously at slot ${oldLayout[slotNumber].slot} has incorrectly changed size from ${oldGapSize} to ${newGapSize} in new layout. Should be ${oldGapSize - newSlots}`);
      }
    }
  }

  // Check that gaps are aligned to the 50th slot in the new layout

  // Find last gap slot
  let lastGapSlot = newLayout.length;
  for(let i = newLayout.length - 1; i >= 0; i--) {
    if(newLayout[i].name === '__gap') {
      lastGapSlot = i;
      break;
    }
  }

  for(let i = 0; i < lastGapSlot; i++) { // Ignore last gap
    if(newLayout[i].name === '__gap') {
      // Find the next non-gap slot
      let nextDirtySlot = i;
      while(newLayout[nextDirtySlot].empty && nextDirtySlot < newLayout.length - 1) {
        nextDirtySlot++;
      }

      const nextSlot = newLayout[nextDirtySlot].slot;
      if(nextSlot % 50 !== 1) {
        warnings.push(`Next slot after the gap at ${newLayout[i].slot} is not aligned to the x*50th + 1 slot`);
      }
    }
  }

  // Sanity check that all slots in old layout have been read
  for (let i = 0; i < oldLayout.length; i++) {
    if (!oldLayout[i].read) {
      errors.push(new Error(`Slot with name ${oldLayout[i].name} of type ${oldLayout[i].type} has been removed`));
    }
  }

  // Delete files if keep is not specified
  if (!args.keep) {
    fs.unlinkSync(args.old as string);
    fs.unlinkSync(args.new as string);
  }
  
  // Print warnings
  if (warnings.length > 0) {
    for (let i = 0; i < warnings.length; i++) {
      console.log(chalk.yellow(warnings[i]));
    }
  }

  // Print errors or success
  if (errors.length > 0) {
    const aggregatedError = new Error(chalk.red("Errors found in storage layout"));
    (aggregatedError as any).error = errors;
    throw aggregatedError;
  } else {
    console.log(chalk.bold.green('Storage layout is upgrade safe'));
  }
}

function formatCSV(csv: string) {
  const layoutFormatted = csv.split('\n')
  // Array with all storage slots
  const storageSlots: StorageSlot[] = [];

  // Begin iterating when table begins, ignore any output before
  let storageBegin = 2; // storage begins on line 2 if there is no extra output
  for (let i = 0; i < layoutFormatted.length; i++) {
    if (layoutFormatted[i].includes('| Name')){
      break;
    }
    storageBegin++;
  }

  for (let i = storageBegin; i < layoutFormatted.length; i++) {
    const data = layoutFormatted[i].split('|'); 
    
    // Reached EOF
    if (data.length <= 1) {
      break;
    }

    const isEmpty = data[1].trim() === '__gap' ? true : false;
    // Extract the relevant data from each line
    const entry = {
      name: data[1].trim(),
      type: data[2].trim(),
      slot: Number(data[3].trim()),
      offset: Number(data[4].trim()),
      bytes: Number(data[5].trim()),
      empty: isEmpty,
      read: false
    }

    // Push the entry to the storageSlots array
    storageSlots.push(entry);

    if (entry.name == '__gap'){
      // Push empty slots for size of gap
      const oldGapSize = getGapSize(entry.type);
      for (let j = 1; j < oldGapSize; j++) {
        storageSlots.push({
          name: '',
          type: '',
          slot: entry.slot + j,
          offset: 0,
          bytes: 0,
          empty: true,
          read: false
        });
      }
    }
  }

  // Assert that the storageSlots array is sorted by slot
  for (let i = 0; i < storageSlots.length - 1; i++) {
    if (storageSlots[i].slot > storageSlots[i+1].slot) {
      throw new Error('Storage slots are not sorted by slot number');
    }
  }

  return storageSlots;
}

// Gap in format `uint256[x]`, return x
function getGapSize(gap: string): number {
  return Number(gap.split('[')[1].split(']')[0]);
}

validateStorageSlots();


type StorageSlot = {
  name: string,
  type: string,
  slot: number,
  offset: number,
  bytes: number,
  empty: boolean,
  read: boolean
}