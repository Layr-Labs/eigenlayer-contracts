import "../ERC20/erc20cvl.spec";
import "../ERC20/WETHcvl.spec";

import "../problems.spec";
import "../unresolved.spec";
import "../optimizations.spec";

use builtin rule sanity filtered { f -> f.contract == currentContract }
