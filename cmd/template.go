package cmd

const Template = `
require 'fofa_core'

class FofaExploits < Fofa::Exploit
  def get_info
    {
        'Name': 'Exploit Name',
        'Description': 'Exploit Description',
        'Author': 'Author',
        'Product': 'Product Name',
        'Homepage': 'https://fofa.so/',
        'References':
            [
                "https://fofa.so/"
            ],
        'DisclosureDate': '2001-01-01',
        'FofaQuery':'domain="fofa.so"',
        'ScanSteps':
            [
                {
                    'Request':
                        {
                            method: 'GET',
                            uri: '/',
                            header: {},
                            data: nil
                        },
                    'ResponseTest': {
                        type: 'group',
                        operation: 'AND',
                        checks: [
                            {
                                type: 'item',
                                variable: '$code',
                                operation: '==',
                                value: 200
                            },
                            {
                                type: 'item',
                                variable: '$body',
                                operation: 'regex',
                                value: 'test'
                            }
                        ]
                    }
                },
            ]
    }
  end

  def initialize(info = {})
    super( info.merge(get_info()) )
  end

  def vulnerable(hostinfo)
    excute_scansteps(hostinfo) if @info['ScanSteps']
  end

  def exploit(hostinfo)
  end
end
`
