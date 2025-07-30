
export enum ResultType {
    Ok = "Ok",
    Warnning = "Warnning",
    Error = "Error",
}

export type ResponseModel = {
    result: ResultType;
    messages: [];
    data: {};
};