import "../ERC20/erc20dispatched.spec";

import "../problems.spec";
import "../optimizations.spec";
import "../eigenpod.spec";

use builtin rule sanity filtered { f -> f.contract == currentContract }
