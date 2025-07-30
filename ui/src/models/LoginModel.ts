
export type LoginModel = {
    email: string;
    password: string;
    remmember: boolean;
};


export const InitLoginModel: LoginModel = {
    email: '',
    password: '',
    remmember: false,
};
