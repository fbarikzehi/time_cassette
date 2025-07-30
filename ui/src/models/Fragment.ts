
export interface FragmentModel {
    Id: string,
    name: string,
    color: string,
    status: false,
    cassetteName: string,
    cassetteId: string,
    counts: FragmentCountModel
}
export interface FragmentCountModel {
    branch: number,
}
export type CreateModel = {
    name: string,
    cassetteId: string,
}

