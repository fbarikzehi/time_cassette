
export interface BranchModel {
    Id: string,
    name: string,
    color: string,
    status: false,
    fragmentName: string,
    fragmentId: string,
    counts: BranchCountModel,
    userName: string
}

export interface BranchCountModel {
    time: number,
}

export type CreateModel = {
    name: string,
    fragmentId: string,
    handlerUserId: string,
}

export type UpdateModel = {
    id: string,
    name: string,
    fragmentId: string,
    handlerUserId: string,
}

