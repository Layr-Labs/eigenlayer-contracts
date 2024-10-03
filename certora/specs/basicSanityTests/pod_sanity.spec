import "../problems.spec";
import "../eigenpod.spec";
import "../optimizations.spec";

use builtin rule sanity filtered { f -> f.contract == currentContract }
