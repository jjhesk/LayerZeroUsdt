/* DO NOT EDIT THE BELOW FILE AS THIS IS LIKELY WILL BE GENERATED AGAIN AND REWRITE OVER IT */

// tslint:disable:no-consecutive-blank-lines ordered-imports align trailing-comma enum-naming
// tslint:disable:whitespace no-unbound-method no-trailing-whitespace
// tslint:disable:no-unused-variable

import * as ethers from "ethers"
// eslint-disable-next-line import/named
import {
  assert,
  schemas,
  // eslint-disable-next-line import/named
  SubscriptionManager,
  // eslint-disable-next-line import/named
  BaseContract,
  // eslint-disable-next-line import/named
  EventCallback,
  // eslint-disable-next-line import/named
  IndexedFilterValues,
  // eslint-disable-next-line import/named
  BlockRange,
  // eslint-disable-next-line import/named
  DecodedLogArgs,
  // eslint-disable-next-line import/named
  LogWithDecodedArgs,
  // eslint-disable-next-line import/named
  MethodAbi
} from "vue-blocklink"

import {
  BatchRequest,
  Extension,
  Log,
  PromiEvent,
  provider,
  Providers,
  RLPEncodedTransaction,
  Transaction,
  TransactionConfig,
  TransactionReceipt,
  Common,
  hardfork,
  chain,
  BlockNumber,
  LogsOptions,
  PastLogsOptions
} from "web3-core"

import { Utils, AbiItem } from "web3-utils"
import Web3 from "web3"
import BN from "BN.js"

// tslint:enable:no-unused-variable
export interface ContractInterface {
    DOMAIN_TYPEHASH():Promise<string>
    PERMIT_TYPEHASH():Promise<string>
    addBlackList(evil: string,):Promise<void>
    allowance(owner: string,spender: string,):Promise<BN>
    approve(spender: string,amount: BN,):Promise<boolean>
    balanceOf(account: string,):Promise<BN>
    decimals():Promise<BN>
    decreaseAllowance(spender: string,subtractedValue: BN,):Promise<boolean>
    destroyBlackFunds(evil: string,):Promise<void>
    endpoint():Promise<string>
    getBlackListStatus(_maker: string,):Promise<boolean>
    governance():Promise<string>
    increaseAllowance(spender: string,addedValue: BN,):Promise<boolean>
    isBlackListed(index_0: string,):Promise<boolean>
    lzReceive(_srcChainId: number|BN,_srcAddress: string,index_2: BN,_payload: string,):Promise<void>
    minters(index_0: string,):Promise<boolean>
    name():Promise<string>
    nonces(index_0: string,):Promise<BN>
    permit(owner: string,spender: string,rawAmount: BN,deadline: BN,v: number|BN,r: string,s: string,):Promise<void>
    remotes(index_0: number|BN,):Promise<string>
    removeBlackList(noevil: string,):Promise<void>
    sendTokens(_chainId: number|BN,_dstOmniChainTokenAddr: string,_qty: BN,coin: BN):Promise<void>
    setGovernance(_governance: string,):Promise<void>
    setRemote(_chainId: number|BN,_remoteAddress: string,):Promise<void>
    symbol():Promise<string>
    totalSupply():Promise<BN>
    transfer(recipient: string,amount: BN,):Promise<boolean>
    transferFrom(sender: string,recipient: string,amount: BN,):Promise<boolean>
}





export enum LayerZeroDollarEvents {
    AddedBlackList = 'AddedBlackList',
    Approval = 'Approval',
    DestroyedBlackFunds = 'DestroyedBlackFunds',
    RemovedBlackList = 'RemovedBlackList',
    Transfer = 'Transfer',
}

export interface LayerZeroDollarAddedBlackListEventArgs extends DecodedLogArgs {
    _user: string;
}

export interface LayerZeroDollarApprovalEventArgs extends DecodedLogArgs {
    owner: string;
    spender: string;
    value: BN;
}

export interface LayerZeroDollarDestroyedBlackFundsEventArgs extends DecodedLogArgs {
    _blackListedUser: string;
    _balance: BN;
}

export interface LayerZeroDollarRemovedBlackListEventArgs extends DecodedLogArgs {
    _user: string;
}

export interface LayerZeroDollarTransferEventArgs extends DecodedLogArgs {
    from: string;
    to: string;
    value: BN;
}


export type LayerZeroDollarEventArgs =
    | LayerZeroDollarAddedBlackListEventArgs
    | LayerZeroDollarApprovalEventArgs
    | LayerZeroDollarDestroyedBlackFundsEventArgs
    | LayerZeroDollarRemovedBlackListEventArgs
    | LayerZeroDollarTransferEventArgs;




/* istanbul ignore next */
// tslint:disable:array-type
// tslint:disable:no-parameter-reassignment
// tslint:disable-next-line:class-name
export class LayerZeroDollarContract extends BaseContract implements ContractInterface{
    /**
     * @ignore
     */
public static deployedBytecode: string | undefined;
public static readonly contractName = 'LayerZeroDollar';
    private readonly _methodABIIndex: { [name: string]: number } = {};
    //todo: we will come back and fix it later not generic Error @https://www.typescriptlang.org/docs/handbook/2/conditional-types.html
    // @ts-ignore
    private readonly _subscriptionManager: SubscriptionManager<LayerZeroDollarEventArgs, LayerZeroDollarEvents>;


    public static Instance(): (LayerZeroDollarContract | any | boolean) {
        if (window && window.hasOwnProperty("__eth_contract_LayerZeroDollar")) {
          // @ts-ignore
          const obj = window.__eth_contract_LayerZeroDollar
          if (obj instanceof LayerZeroDollarContract) {
            return (obj) as LayerZeroDollarContract
          } else {
            return (obj) as LayerZeroDollarContract
          }
        } else {
          return false
        }
    }

    static async init(
        contract_address: string,
        supportedProvider: provider,
        ww3: Web3
        ):Promise<LayerZeroDollarContract>
    {
        const contractInstance = await new LayerZeroDollarContract(
            contract_address,
            supportedProvider,
            ww3
        );

        contractInstance.constructorArgs = [/** "_layerZeroEndpoint"
 **/];

        if (window && !window.hasOwnProperty("__eth_contract_LayerZeroDollar")) {
            // @ts-ignore
            window.__eth_contract_LayerZeroDollar = contractInstance
        }

        return contractInstance
    }

    /**
     * @returns The contract ABI
     */
    public static ABI(): AbiItem[] {
        const abi = [
            { 
                inputs: [
                    {
                        name: '_layerZeroEndpoint',
                        type: 'address',
                    },
                ],
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'constructor',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: '_user',
                        type: 'address',
                        indexed: false,
                    },
                ],
                name: 'AddedBlackList',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'owner',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'spender',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'value',
                        type: 'uint256',
                        indexed: false,
                    },
                ],
                name: 'Approval',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: '_blackListedUser',
                        type: 'address',
                        indexed: false,
                    },
                    {
                        name: '_balance',
                        type: 'uint256',
                        indexed: false,
                    },
                ],
                name: 'DestroyedBlackFunds',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: '_user',
                        type: 'address',
                        indexed: false,
                    },
                ],
                name: 'RemovedBlackList',
                outputs: [
                ],
                type: 'event',
            },
            { 
                anonymous: false,
                inputs: [
                    {
                        name: 'from',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'to',
                        type: 'address',
                        indexed: true,
                    },
                    {
                        name: 'value',
                        type: 'uint256',
                        indexed: false,
                    },
                ],
                name: 'Transfer',
                outputs: [
                ],
                type: 'event',
            },
            { 
                inputs: [
                ],
                name: 'DOMAIN_TYPEHASH',
                outputs: [
                    {
                        name: '',
                        type: 'bytes32',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'PERMIT_TYPEHASH',
                outputs: [
                    {
                        name: '',
                        type: 'bytes32',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'evil',
                        type: 'address',
                    },
                ],
                name: 'addBlackList',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'owner',
                        type: 'address',
                    },
                    {
                        name: 'spender',
                        type: 'address',
                    },
                ],
                name: 'allowance',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'spender',
                        type: 'address',
                    },
                    {
                        name: 'amount',
                        type: 'uint256',
                    },
                ],
                name: 'approve',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'account',
                        type: 'address',
                    },
                ],
                name: 'balanceOf',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'decimals',
                outputs: [
                    {
                        name: '',
                        type: 'uint8',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'spender',
                        type: 'address',
                    },
                    {
                        name: 'subtractedValue',
                        type: 'uint256',
                    },
                ],
                name: 'decreaseAllowance',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'evil',
                        type: 'address',
                    },
                ],
                name: 'destroyBlackFunds',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'endpoint',
                outputs: [
                    {
                        name: '',
                        type: 'address',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: '_maker',
                        type: 'address',
                    },
                ],
                name: 'getBlackListStatus',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'governance',
                outputs: [
                    {
                        name: '',
                        type: 'address',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'spender',
                        type: 'address',
                    },
                    {
                        name: 'addedValue',
                        type: 'uint256',
                    },
                ],
                name: 'increaseAllowance',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'index_0',
                        type: 'address',
                    },
                ],
                name: 'isBlackListed',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: '_srcChainId',
                        type: 'uint16',
                    },
                    {
                        name: '_srcAddress',
                        type: 'bytes',
                    },
                    {
                        name: 'index_2',
                        type: 'uint64',
                    },
                    {
                        name: '_payload',
                        type: 'bytes',
                    },
                ],
                name: 'lzReceive',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'index_0',
                        type: 'address',
                    },
                ],
                name: 'minters',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'name',
                outputs: [
                    {
                        name: '',
                        type: 'string',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'index_0',
                        type: 'address',
                    },
                ],
                name: 'nonces',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'owner',
                        type: 'address',
                    },
                    {
                        name: 'spender',
                        type: 'address',
                    },
                    {
                        name: 'rawAmount',
                        type: 'uint256',
                    },
                    {
                        name: 'deadline',
                        type: 'uint256',
                    },
                    {
                        name: 'v',
                        type: 'uint8',
                    },
                    {
                        name: 'r',
                        type: 'bytes32',
                    },
                    {
                        name: 's',
                        type: 'bytes32',
                    },
                ],
                name: 'permit',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'index_0',
                        type: 'uint16',
                    },
                ],
                name: 'remotes',
                outputs: [
                    {
                        name: '',
                        type: 'bytes',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'noevil',
                        type: 'address',
                    },
                ],
                name: 'removeBlackList',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: '_chainId',
                        type: 'uint16',
                    },
                    {
                        name: '_dstOmniChainTokenAddr',
                        type: 'bytes',
                    },
                    {
                        name: '_qty',
                        type: 'uint256',
                    },
                ],
                name: 'sendTokens',
                outputs: [
                ],
                stateMutability: 'payable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: '_governance',
                        type: 'address',
                    },
                ],
                name: 'setGovernance',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: '_chainId',
                        type: 'uint16',
                    },
                    {
                        name: '_remoteAddress',
                        type: 'bytes',
                    },
                ],
                name: 'setRemote',
                outputs: [
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'symbol',
                outputs: [
                    {
                        name: '',
                        type: 'string',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                ],
                name: 'totalSupply',
                outputs: [
                    {
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'recipient',
                        type: 'address',
                    },
                    {
                        name: 'amount',
                        type: 'uint256',
                    },
                ],
                name: 'transfer',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
            { 
                inputs: [
                    {
                        name: 'sender',
                        type: 'address',
                    },
                    {
                        name: 'recipient',
                        type: 'address',
                    },
                    {
                        name: 'amount',
                        type: 'uint256',
                    },
                ],
                name: 'transferFrom',
                outputs: [
                    {
                        name: '',
                        type: 'bool',
                    },
                ],
                stateMutability: 'nonpayable',
                type: 'function',
            },
        ] as AbiItem[];
        return abi;
    }

    /**
     the listed content for the contract functions
    **/

    public async DOMAIN_TYPEHASH(): Promise<string> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.DOMAIN_TYPEHASH(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async DOMAIN_TYPEHASHGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.DOMAIN_TYPEHASH().estimateGas();
        return gasAmount;
    };


    public async PERMIT_TYPEHASH(): Promise<string> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.PERMIT_TYPEHASH(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async PERMIT_TYPEHASHGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.PERMIT_TYPEHASH().estimateGas();
        return gasAmount;
    };


    public async addBlackList(evil: string,): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('evil', evil);

        const promizz = self._contract.methods.addBlackList(
            evil,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async addBlackListGas(evil: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.addBlackList(evil,).estimateGas();
        return gasAmount;
    };


    public async allowance(owner: string,spender: string,): Promise<BN> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('owner', owner);
            assert.isString('spender', spender);

        const promizz = self._contract.methods.allowance(
            owner,
                    spender,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async allowanceGas(owner: string,spender: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.allowance(owner,spender,).estimateGas();
        return gasAmount;
    };


    public async approve(spender: string,amount: BN,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('spender', spender);
            assert.isNumberOrBigNumber('amount', amount);

        const promizz = self._contract.methods.approve(
            spender,
                    amount,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async approveGas(spender: string,amount: BN,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.approve(spender,amount,).estimateGas();
        return gasAmount;
    };


    public async balanceOf(account: string,): Promise<BN> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('account', account);

        const promizz = self._contract.methods.balanceOf(
            account,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async balanceOfGas(account: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.balanceOf(account,).estimateGas();
        return gasAmount;
    };


    public async decimals(): Promise<BN> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.decimals(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async decimalsGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.decimals().estimateGas();
        return gasAmount;
    };


    public async decreaseAllowance(spender: string,subtractedValue: BN,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('spender', spender);
            assert.isNumberOrBigNumber('subtractedValue', subtractedValue);

        const promizz = self._contract.methods.decreaseAllowance(
            spender,
                    subtractedValue,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async decreaseAllowanceGas(spender: string,subtractedValue: BN,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.decreaseAllowance(spender,subtractedValue,).estimateGas();
        return gasAmount;
    };


    public async destroyBlackFunds(evil: string,): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('evil', evil);

        const promizz = self._contract.methods.destroyBlackFunds(
            evil,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async destroyBlackFundsGas(evil: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.destroyBlackFunds(evil,).estimateGas();
        return gasAmount;
    };


    public async endpoint(): Promise<string> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.endpoint(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async endpointGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.endpoint().estimateGas();
        return gasAmount;
    };


    public async getBlackListStatus(_maker: string,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('_maker', _maker);

        const promizz = self._contract.methods.getBlackListStatus(
            _maker,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async getBlackListStatusGas(_maker: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.getBlackListStatus(_maker,).estimateGas();
        return gasAmount;
    };


    public async governance(): Promise<string> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.governance(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async governanceGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.governance().estimateGas();
        return gasAmount;
    };


    public async increaseAllowance(spender: string,addedValue: BN,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('spender', spender);
            assert.isNumberOrBigNumber('addedValue', addedValue);

        const promizz = self._contract.methods.increaseAllowance(
            spender,
                    addedValue,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async increaseAllowanceGas(spender: string,addedValue: BN,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.increaseAllowance(spender,addedValue,).estimateGas();
        return gasAmount;
    };


    public async isBlackListed(index_0: string,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('index_0', index_0);

        const promizz = self._contract.methods.isBlackListed(
            index_0,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async isBlackListedGas(index_0: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.isBlackListed(index_0,).estimateGas();
        return gasAmount;
    };


    public async lzReceive(_srcChainId: number|BN,_srcAddress: string,index_2: BN,_payload: string,): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isNumberOrBigNumber('_srcChainId', _srcChainId);
            assert.isString('_srcAddress', _srcAddress);
            assert.isBigNumber('index_2', index_2);
            assert.isString('_payload', _payload);

        const promizz = self._contract.methods.lzReceive(
            _srcChainId,
                    _srcAddress,
                    index_2,
                    _payload,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async lzReceiveGas(_srcChainId: number|BN,_srcAddress: string,index_2: BN,_payload: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.lzReceive(_srcChainId,_srcAddress,index_2,_payload,).estimateGas();
        return gasAmount;
    };


    public async minters(index_0: string,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('index_0', index_0);

        const promizz = self._contract.methods.minters(
            index_0,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async mintersGas(index_0: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.minters(index_0,).estimateGas();
        return gasAmount;
    };


    public async name(): Promise<string> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.name(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async nameGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.name().estimateGas();
        return gasAmount;
    };


    public async nonces(index_0: string,): Promise<BN> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('index_0', index_0);

        const promizz = self._contract.methods.nonces(
            index_0,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async noncesGas(index_0: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.nonces(index_0,).estimateGas();
        return gasAmount;
    };


    public async permit(owner: string,spender: string,rawAmount: BN,deadline: BN,v: number|BN,r: string,s: string,): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('owner', owner);
            assert.isString('spender', spender);
            assert.isNumberOrBigNumber('rawAmount', rawAmount);
            assert.isNumberOrBigNumber('deadline', deadline);
            assert.isNumberOrBigNumber('v', v);
            assert.isString('r', r);
            assert.isString('s', s);

        const promizz = self._contract.methods.permit(
            owner,
                    spender,
                    rawAmount,
                    deadline,
                    v,
                    r,
                    s,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async permitGas(owner: string,spender: string,rawAmount: BN,deadline: BN,v: number|BN,r: string,s: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.permit(owner,spender,rawAmount,deadline,v,r,s,).estimateGas();
        return gasAmount;
    };


    public async remotes(index_0: number|BN,): Promise<string> {
        const self = this as any as LayerZeroDollarContract;

            assert.isNumberOrBigNumber('index_0', index_0);

        const promizz = self._contract.methods.remotes(
            index_0,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async remotesGas(index_0: number|BN,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.remotes(index_0,).estimateGas();
        return gasAmount;
    };


    public async removeBlackList(noevil: string,): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('noevil', noevil);

        const promizz = self._contract.methods.removeBlackList(
            noevil,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async removeBlackListGas(noevil: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.removeBlackList(noevil,).estimateGas();
        return gasAmount;
    };


    public async sendTokens(_chainId: number|BN,_dstOmniChainTokenAddr: string,_qty: BN,valCoin: BN): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isNumberOrBigNumber('_chainId', _chainId);
            assert.isString('_dstOmniChainTokenAddr', _dstOmniChainTokenAddr);
            assert.isNumberOrBigNumber('_qty', _qty);

        const promizz = self._contract.methods.sendTokens(
            _chainId,
                    _dstOmniChainTokenAddr,
                    _qty,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice, value: valCoin
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async sendTokensGas(_chainId: number|BN,_dstOmniChainTokenAddr: string,_qty: BN,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.sendTokens(_chainId,_dstOmniChainTokenAddr,_qty,).estimateGas();
        return gasAmount;
    };


    public async setGovernance(_governance: string,): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('_governance', _governance);

        const promizz = self._contract.methods.setGovernance(
            _governance,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async setGovernanceGas(_governance: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.setGovernance(_governance,).estimateGas();
        return gasAmount;
    };


    public async setRemote(_chainId: number|BN,_remoteAddress: string,): Promise<void> {
        const self = this as any as LayerZeroDollarContract;

            assert.isNumberOrBigNumber('_chainId', _chainId);
            assert.isString('_remoteAddress', _remoteAddress);

        const promizz = self._contract.methods.setRemote(
            _chainId,
                    _remoteAddress,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async setRemoteGas(_chainId: number|BN,_remoteAddress: string,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.setRemote(_chainId,_remoteAddress,).estimateGas();
        return gasAmount;
    };


    public async symbol(): Promise<string> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.symbol(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async symbolGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.symbol().estimateGas();
        return gasAmount;
    };


    public async totalSupply(): Promise<BN> {
        const self = this as any as LayerZeroDollarContract;


        const promizz = self._contract.methods.totalSupply(
)

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


        const result = await promizz.call(_essence);

        return result;
    };


    public async totalSupplyGas(): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.totalSupply().estimateGas();
        return gasAmount;
    };


    public async transfer(recipient: string,amount: BN,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('recipient', recipient);
            assert.isNumberOrBigNumber('amount', amount);

        const promizz = self._contract.methods.transfer(
            recipient,
                    amount,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async transferGas(recipient: string,amount: BN,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.transfer(recipient,amount,).estimateGas();
        return gasAmount;
    };


    public async transferFrom(sender: string,recipient: string,amount: BN,): Promise<boolean> {
        const self = this as any as LayerZeroDollarContract;

            assert.isString('sender', sender);
            assert.isString('recipient', recipient);
            assert.isNumberOrBigNumber('amount', amount);

        const promizz = self._contract.methods.transferFrom(
            sender,
                    recipient,
                    amount,
        )

        const _essence = {
            from: this._blockwrap.getAccountAddress(),
            gas: this.gas,
            gasPrice: this.gasPrice
        };


const result = await promizz.send(_essence)
        .on('transactionHash', (hash) => {
            this.onBroadcast(hash);
        }).on('confirmation', (confirmationNumber, receipt) => {
            this.onConfirmation(receipt);
        }).on('receipt', (r) => {
            this.pushReceiptSuccess(r);
        }).on('error', (error, receipt) => {
            this.onError(receipt, error);
        }).catch((e) => {
            this.catchErro(e)
        });

        return result;
    };


    public async transferFromGas(sender: string,recipient: string,amount: BN,): Promise<number>{
        const self = this as any as LayerZeroDollarContract;
        const gasAmount = await self._contract.methods.transferFrom(sender,recipient,amount,).estimateGas();
        return gasAmount;
    };


    /**
     * Subscribe to an event type emitted by the LayerZeroDollar contract.
     * @param eventName The LayerZeroDollar contract event you would like to subscribe to.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{maker: aUserAddressHex}`
     * @param callback Callback that gets called when a log is added/removed
     * @param isVerbose Enable verbose subscription warnings (e.g recoverable network issues encountered)
     * @return Subscription token used later to unsubscribe
     */
    public subscribe<ArgsType extends LayerZeroDollarEventArgs>(
        eventName: LayerZeroDollarEvents,
        indexFilterValues: IndexedFilterValues,
        callback: EventCallback<ArgsType>,
        isVerbose: boolean = false,
        blockPollingIntervalMs?: number,
    ): string {
        assert.doesBelongToStringEnum('eventName', eventName, LayerZeroDollarEvents);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        assert.isFunction('callback', callback);
        const subscriptionToken = this._subscriptionManager.subscribe<ArgsType>(
            this._address,
            eventName,
            indexFilterValues,
            LayerZeroDollarContract.ABI(),
            callback,
            isVerbose,
            blockPollingIntervalMs,
        );
        return subscriptionToken;
    }

    /**
     * Cancel a subscription
     * @param subscriptionToken Subscription token returned by `subscribe()`
     */
    public unsubscribe(subscriptionToken: string): void {
        this._subscriptionManager.unsubscribe(subscriptionToken);
    }

    /**
     * Cancels all existing subscriptions
     */
    public unsubscribeAll(): void {
        this._subscriptionManager.unsubscribeAll();
    }


    /**
     * Gets historical logs without creating a subscription
     * @param eventName The LayerZeroDollar contract event you would like to subscribe to.
     * @param blockRange Block range to get logs from.
     * @param indexFilterValues An object where the keys are indexed args returned by the event and
     * the value is the value you are interested in. E.g `{_from: aUserAddressHex}`
     * @return Array of logs that match the parameters
     */
    public async getLogsAsync<ArgsType extends LayerZeroDollarEventArgs>(
        eventName: LayerZeroDollarEvents,
        blockRange: BlockRange,
        indexFilterValues: IndexedFilterValues,
    ): Promise<Array<LogWithDecodedArgs<ArgsType>>> {
        assert.doesBelongToStringEnum('eventName', eventName, LayerZeroDollarEvents);
        assert.doesConformToSchema('blockRange', blockRange, schemas.blockRangeSchema);
        assert.doesConformToSchema('indexFilterValues', indexFilterValues, schemas.indexFilterValuesSchema);
        const logs = await this._subscriptionManager.getLogsAsync<ArgsType>(
            this._address,
            eventName,
            blockRange,
            indexFilterValues,
            LayerZeroDollarContract.ABI(),
        );
        return logs;
    }

    constructor(
        address: string,
        supportedProvider: provider,
        ww3: Web3
    ) {
        super('LayerZeroDollar', LayerZeroDollarContract.ABI(), address, supportedProvider,ww3);
        this._subscriptionManager = new SubscriptionManager(
            LayerZeroDollarContract.ABI(),
            supportedProvider,
        );

        LayerZeroDollarContract.ABI().forEach((item, index) => {
            if (item.type === 'function') {
                const methodAbi = item as MethodAbi;
                this._methodABIIndex[methodAbi.name] = index;
            }
        });


    }
}

// tslint:disable:max-file-line-count
// tslint:enable:no-unbound-method no-parameter-reassignment no-consecutive-blank-lines ordered-imports align
// tslint:enable:trailing-comma whitespace no-trailing-whitespace
