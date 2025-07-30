
export interface TimeModel {
    Id: string,
    name: string,
    color: string,
    startDateTime: string,
    endDateTime: string,
    branchName: string,
    branchId: string,
}

export type CreateModel = {
    duration: string,
    startPointDateTime: string,
    branchId: string,
}

export type UpdateModel = {
    id: string,
    duration: string,
    startPointDateTime: string,
    branchId: string,
}

