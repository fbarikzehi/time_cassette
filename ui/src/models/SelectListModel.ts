

export interface SeletcListOption {
    name: string,
    value: string,
    selected: boolean
}

export interface SeletcList {
    id: string,
    name: string,
    options: Array<SeletcListOption>
}