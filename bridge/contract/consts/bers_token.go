package consts

const BersTokenABI = `[
  {
    "inputs": [
      {
        "internalType": "uint256",
        "name": "_totalSupply",
        "type": "uint256"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "owner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      }
    ],
    "name": "Approval",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "previousOwner",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "OwnershipTransferred",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "account",
        "type": "address"
      }
    ],
    "name": "Paused",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "indexed": true,
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "indexed": false,
        "internalType": "uint256",
        "name": "value",
        "type": "uint256"
      }
    ],
    "name": "Transfer",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "internalType": "address",
        "name": "account",
        "type": "address"
      }
    ],
    "name": "Unpaused",
    "type": "event"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "owner",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      }
    ],
    "name": "allowance",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "approve",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "account",
        "type": "address"
      }
    ],
    "name": "balanceOf",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "decimals",
    "outputs": [
      {
        "internalType": "uint8",
        "name": "",
        "type": "uint8"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "subtractedValue",
        "type": "uint256"
      }
    ],
    "name": "decreaseAllowance",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "spender",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "addedValue",
        "type": "uint256"
      }
    ],
    "name": "increaseAllowance",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "name",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "owner",
    "outputs": [
      {
        "internalType": "address",
        "name": "",
        "type": "address"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "pause",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "paused",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "renounceOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "symbol",
    "outputs": [
      {
        "internalType": "string",
        "name": "",
        "type": "string"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "totalSupply",
    "outputs": [
      {
        "internalType": "uint256",
        "name": "",
        "type": "uint256"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "transfer",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "from",
        "type": "address"
      },
      {
        "internalType": "address",
        "name": "to",
        "type": "address"
      },
      {
        "internalType": "uint256",
        "name": "amount",
        "type": "uint256"
      }
    ],
    "name": "transferFrom",
    "outputs": [
      {
        "internalType": "bool",
        "name": "",
        "type": "bool"
      }
    ],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "internalType": "address",
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "transferOwnership",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "unpause",
    "outputs": [],
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`
const BersTokenBin = "0x60806040523480156200001157600080fd5b50604051620020ff380380620020ff8339818101604052810190620000379190620004be565b6040518060400160405280600881526020017f42657273537761700000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4253500000000000000000000000000000000000000000000000000000000000815250620000c3620000b76200012b60201b60201c565b6200013360201b60201c565b8160049080519060200190620000db929190620003f7565b508060059080519060200190620000f4929190620003f7565b5050506000600660006101000a81548160ff021916908315150217905550620001243382620001f760201b60201c565b5062000748565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200026a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002619062000549565b60405180910390fd5b6200027e600083836200036660201b60201c565b8060036000828254620002929190620005bb565b9250508190555080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200034691906200058d565b60405180910390a36200036260008383620003d660201b60201c565b5050565b6200037e838383620003db60201b620007741760201c565b6200038e620003e060201b60201c565b15620003d1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003c8906200056b565b60405180910390fd5b505050565b505050565b505050565b6000600660009054906101000a900460ff16905090565b828054620004059062000622565b90600052602060002090601f01602090048101928262000429576000855562000475565b82601f106200044457805160ff191683800117855562000475565b8280016001018555821562000475579182015b828111156200047457825182559160200191906001019062000457565b5b50905062000484919062000488565b5090565b5b80821115620004a357600081600090555060010162000489565b5090565b600081519050620004b8816200072e565b92915050565b600060208284031215620004d157600080fd5b6000620004e184828501620004a7565b91505092915050565b6000620004f9601f83620005aa565b91506200050682620006b6565b602082019050919050565b600062000520602a83620005aa565b91506200052d82620006df565b604082019050919050565b620005438162000618565b82525050565b600060208201905081810360008301526200056481620004ea565b9050919050565b60006020820190508181036000830152620005868162000511565b9050919050565b6000602082019050620005a4600083018462000538565b92915050565b600082825260208201905092915050565b6000620005c88262000618565b9150620005d58362000618565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156200060d576200060c62000658565b5b828201905092915050565b6000819050919050565b600060028204905060018216806200063b57607f821691505b6020821081141562000652576200065162000687565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b7f45524332305061757361626c653a20746f6b656e207472616e7366657220776860008201527f696c652070617573656400000000000000000000000000000000000000000000602082015250565b620007398162000618565b81146200074557600080fd5b50565b6119a780620007586000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806370a08231116100a257806395d89b411161007157806395d89b4114610284578063a457c2d7146102a2578063a9059cbb146102d2578063dd62ed3e14610302578063f2fde38b146103325761010b565b806370a0823114610222578063715018a6146102525780638456cb591461025c5780638da5cb5b146102665761010b565b8063313ce567116100de578063313ce567146101ac57806339509351146101ca5780633f4ba83a146101fa5780635c975abb146102045761010b565b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015e57806323b872dd1461017c575b600080fd5b61011861034e565b60405161012591906112b4565b60405180910390f35b61014860048036038101906101439190611029565b6103e0565b6040516101559190611299565b60405180910390f35b610166610403565b6040516101739190611456565b60405180910390f35b61019660048036038101906101919190610fda565b61040d565b6040516101a39190611299565b60405180910390f35b6101b461043c565b6040516101c19190611471565b60405180910390f35b6101e460048036038101906101df9190611029565b610445565b6040516101f19190611299565b60405180910390f35b61020261047c565b005b61020c61048e565b6040516102199190611299565b60405180910390f35b61023c60048036038101906102379190610f75565b6104a5565b6040516102499190611456565b60405180910390f35b61025a6104ee565b005b610264610502565b005b61026e610514565b60405161027b919061127e565b60405180910390f35b61028c61053d565b60405161029991906112b4565b60405180910390f35b6102bc60048036038101906102b79190611029565b6105cf565b6040516102c99190611299565b60405180910390f35b6102ec60048036038101906102e79190611029565b610646565b6040516102f99190611299565b60405180910390f35b61031c60048036038101906103179190610f9e565b610669565b6040516103299190611456565b60405180910390f35b61034c60048036038101906103479190610f75565b6106f0565b005b60606004805461035d90611586565b80601f016020809104026020016040519081016040528092919081815260200182805461038990611586565b80156103d65780601f106103ab576101008083540402835291602001916103d6565b820191906000526020600020905b8154815290600101906020018083116103b957829003601f168201915b5050505050905090565b6000806103eb610779565b90506103f8818585610781565b600191505092915050565b6000600354905090565b600080610418610779565b905061042585828561094c565b6104308585856109d8565b60019150509392505050565b60006012905090565b600080610450610779565b90506104718185856104628589610669565b61046c91906114a8565b610781565b600191505092915050565b610484610c53565b61048c610cd1565b565b6000600660009054906101000a900460ff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6104f6610c53565b6105006000610d34565b565b61050a610c53565b610512610df8565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606005805461054c90611586565b80601f016020809104026020016040519081016040528092919081815260200182805461057890611586565b80156105c55780601f1061059a576101008083540402835291602001916105c5565b820191906000526020600020905b8154815290600101906020018083116105a857829003601f168201915b5050505050905090565b6000806105da610779565b905060006105e88286610669565b90508381101561062d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062490611416565b60405180910390fd5b61063a8286868403610781565b60019250505092915050565b600080610651610779565b905061065e8185856109d8565b600191505092915050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6106f8610c53565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610768576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075f90611316565b60405180910390fd5b61077181610d34565b50565b505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156107f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e8906113f6565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610861576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085890611336565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161093f9190611456565b60405180910390a3505050565b60006109588484610669565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146109d257818110156109c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109bb90611356565b60405180910390fd5b6109d18484848403610781565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610a48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3f906113d6565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ab8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aaf906112d6565b60405180910390fd5b610ac3838383610e5b565b6000600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610b4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4190611376565b60405180910390fd5b818103600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610c3a9190611456565b60405180910390a3610c4d848484610eb3565b50505050565b610c5b610779565b73ffffffffffffffffffffffffffffffffffffffff16610c79610514565b73ffffffffffffffffffffffffffffffffffffffff1614610ccf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc6906113b6565b60405180910390fd5b565b610cd9610eb8565b6000600660006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610d1d610779565b604051610d2a919061127e565b60405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610e00610f01565b6001600660006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610e44610779565b604051610e51919061127e565b60405180910390a1565b610e66838383610774565b610e6e61048e565b15610eae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea590611436565b60405180910390fd5b505050565b505050565b610ec061048e565b610eff576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef6906112f6565b60405180910390fd5b565b610f0961048e565b15610f49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4090611396565b60405180910390fd5b565b600081359050610f5a81611943565b92915050565b600081359050610f6f8161195a565b92915050565b600060208284031215610f8757600080fd5b6000610f9584828501610f4b565b91505092915050565b60008060408385031215610fb157600080fd5b6000610fbf85828601610f4b565b9250506020610fd085828601610f4b565b9150509250929050565b600080600060608486031215610fef57600080fd5b6000610ffd86828701610f4b565b935050602061100e86828701610f4b565b925050604061101f86828701610f60565b9150509250925092565b6000806040838503121561103c57600080fd5b600061104a85828601610f4b565b925050602061105b85828601610f60565b9150509250929050565b61106e816114fe565b82525050565b61107d81611510565b82525050565b600061108e8261148c565b6110988185611497565b93506110a8818560208601611553565b6110b181611616565b840191505092915050565b60006110c9602383611497565b91506110d482611627565b604082019050919050565b60006110ec601483611497565b91506110f782611676565b602082019050919050565b600061110f602683611497565b915061111a8261169f565b604082019050919050565b6000611132602283611497565b915061113d826116ee565b604082019050919050565b6000611155601d83611497565b91506111608261173d565b602082019050919050565b6000611178602683611497565b915061118382611766565b604082019050919050565b600061119b601083611497565b91506111a6826117b5565b602082019050919050565b60006111be602083611497565b91506111c9826117de565b602082019050919050565b60006111e1602583611497565b91506111ec82611807565b604082019050919050565b6000611204602483611497565b915061120f82611856565b604082019050919050565b6000611227602583611497565b9150611232826118a5565b604082019050919050565b600061124a602a83611497565b9150611255826118f4565b604082019050919050565b6112698161153c565b82525050565b61127881611546565b82525050565b60006020820190506112936000830184611065565b92915050565b60006020820190506112ae6000830184611074565b92915050565b600060208201905081810360008301526112ce8184611083565b905092915050565b600060208201905081810360008301526112ef816110bc565b9050919050565b6000602082019050818103600083015261130f816110df565b9050919050565b6000602082019050818103600083015261132f81611102565b9050919050565b6000602082019050818103600083015261134f81611125565b9050919050565b6000602082019050818103600083015261136f81611148565b9050919050565b6000602082019050818103600083015261138f8161116b565b9050919050565b600060208201905081810360008301526113af8161118e565b9050919050565b600060208201905081810360008301526113cf816111b1565b9050919050565b600060208201905081810360008301526113ef816111d4565b9050919050565b6000602082019050818103600083015261140f816111f7565b9050919050565b6000602082019050818103600083015261142f8161121a565b9050919050565b6000602082019050818103600083015261144f8161123d565b9050919050565b600060208201905061146b6000830184611260565b92915050565b6000602082019050611486600083018461126f565b92915050565b600081519050919050565b600082825260208201905092915050565b60006114b38261153c565b91506114be8361153c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156114f3576114f26115b8565b5b828201905092915050565b60006115098261151c565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611571578082015181840152602081019050611556565b83811115611580576000848401525b50505050565b6000600282049050600182168061159e57607f821691505b602082108114156115b2576115b16115e7565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332305061757361626c653a20746f6b656e207472616e7366657220776860008201527f696c652070617573656400000000000000000000000000000000000000000000602082015250565b61194c816114fe565b811461195757600080fd5b50565b6119638161153c565b811461196e57600080fd5b5056fea2646970667358221220866df88433381bd6417bf21126dedbe1fb1ccea7c5aec3c3efcc5056f18982e164736f6c63430008010033"
