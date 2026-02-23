#!/bin/bash

# Apply the patch to the contracts directory
echo "Applying the patch for SplitContractMixin"
git apply certora/patches/SplitContractMixin.patch

echo "Applying the patch for AllocationManager"
git apply certora/patches/AllocationManager.patch