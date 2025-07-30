
export interface CassetteModel {
    Id: string,
    name: string,
    color: string,
    is_private: true,
    status: false,
    counts: CassetteCountModel
}
export interface CassetteCountModel {
    fragment: number,
}
export type CreateModel = {
    name: string
}