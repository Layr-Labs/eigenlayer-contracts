// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "../interfaces/IServiceManager.sol";
import "../interfaces/IQuorumRegistry.sol";
import "../interfaces/IDelegationManager.sol";
import "../interfaces/IPaymentManager.sol";
import "../permissions/Pausable.sol";

/**
 * @title Controls 'rolled-up' middleware payments.
 * @author Layr Labs, Inc.
 * @notice This contract is used for doing interactive payment challenges.
 * @notice The contract is marked as abstract since it does not implement the `respondToPaymentChallengeFinal`
 * function -- see DataLayerPaymentManager for an example
 */
//
abstract contract PaymentManager is Initializable, IPaymentManager, Pausable {
    using SafeERC20 for IERC20;

    uint8 constant internal PAUSED_NEW_PAYMENT_COMMIT = 0;
    uint8 constant internal PAUSED_REDEEM_PAYMENT = 1;

    // DATA STRUCTURES

    /// @notice The ServiceManager contract for this middleware, where tasks are created / initiated.
    IServiceManager public immutable serviceManager;

    /// @notice the ERC20 token that will be used by the disperser to pay the service fees to middleware nodes.
    IERC20 public immutable paymentToken;

    /// @notice Deposits of future fees to be drawn against when paying for service from the middleware
    mapping(address => uint256) public depositsOf;

    /// @notice depositors => addresses approved to spend deposits => allowance
    mapping(address => mapping(address => uint256)) public allowances;

    /// @notice when applied to a function, ensures that the function is only callable by the `serviceManager`
    modifier onlyServiceManager() {
        require(msg.sender == address(serviceManager), "onlyServiceManager");
        _;
    }

    /// @notice when applied to a function, ensures that the function is only callable by the owner of the `serviceManager`
    modifier onlyServiceManagerOwner() {
        require(msg.sender == serviceManager.owner(), "onlyServiceManagerOwner");
        _;
    }

    constructor(
        IServiceManager _serviceManager,
        IERC20 _paymentToken
    ) {
        serviceManager = _serviceManager;
        paymentToken = _paymentToken;
        _disableInitializers();
    }

    function initialize(IPauserRegistry _pauserReg) public initializer {
        _initializePauser(_pauserReg, UNPAUSE_ALL);
    }

    /**
     * @notice deposit one-time fees by the `msg.sender` with this contract to pay for future tasks of this middleware
     * @param depositFor could be the `msg.sender` themselves, or a different address for whom `msg.sender` is depositing these future fees
     * @param amount is amount of futures fees being deposited
     */
    function depositFutureFees(address depositFor, uint256 amount) external {
        paymentToken.safeTransferFrom(msg.sender, address(this), amount);
        depositsOf[depositFor] += amount;
    }

    /// @notice Allows the `allowed` address to spend up to `amount` of the `msg.sender`'s funds that have been deposited in this contract
    function setAllowance(address allowed, uint256 amount) external {
        allowances[msg.sender][allowed] = amount;
    }

    /// @notice Used for deducting the fees from the payer to the middleware
    function takeFee(address initiator, address payer, uint256 feeAmount) external virtual onlyServiceManager {
        if (initiator != payer) {
            if (allowances[payer][initiator] != type(uint256).max) {
                allowances[payer][initiator] -= feeAmount;
            }
        }

        // decrement `payer`'s stored deposits
        depositsOf[payer] -= feeAmount;
    }
}
