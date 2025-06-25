[elip-007]: https://github.com/eigenfoundation/ELIPs/blob/main/ELIPs/ELIP-007.md

## Multichain Docs

The EigenLayer multichain protocol enables *consumption* of EigenLayer L1 stake on supported destination chains.  This document provides an overview of system components, contracts, and user roles and is up-to-date with the latest [ELIP-007][elip-007]. Further documentation on the specific contracts can be found in this folder. 

#### Contents

* [System Diagram](#system-diagram)

### System Diagram

```mermaid
classDiagram 
direction TD
namespace Source Chain {
		class CrossChainRegistry {
				createGenerationReservation
				removeGenerationReservation
				calculateOperatorTableBytes
		}
    class OperatorTableCalculator {
		    ECDSA/BN254 Table Calculator
	  }
		class EigenLayer Core {
				DelegationManager
				AllocationManager
				KeyRegistrar
				PermissionController
		}
}
namespace Off Chain {
    class Generator {
		    generateGlobalTableRoot
		}
		class Transporter {
				confirmGlobalTableRoot
				updateOperatorTable
		}
}
namespace Destination Chain {
    class OperatorTableUpdater {
		    confirmGlobalTableRoot
		    updateOperatorTable
		}
    class ECDSACertificateVerifier {
		    updateOperatorTable
		    verifyCertificate
		}
    class BN254CertificateVerifier {
		    updateOperatorTable
		    verifyCertificate
		}
}


CrossChainRegistry --> OperatorTableCalculator: Reads
OperatorTableCalculator --> EigenLayer Core: Reads
Generator --> CrossChainRegistry: Reads
Transporter --> Generator: Gets Root
Transporter --> OperatorTableUpdater: Confirms Root, Updates Tables
OperatorTableUpdater --> ECDSACertificateVerifier: Updates table
OperatorTableUpdater --> BN254CertificateVerifier: Updates Table
```

### System Contents

#### Cross Chain Registry

#### 