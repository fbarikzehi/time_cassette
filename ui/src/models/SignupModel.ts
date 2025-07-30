
export type SignupModel = {
    email: string;
    password: string;
    confirmPassword: string;
    remmember: boolean;
};


export const InitSignupModel: SignupModel = {
    email: '',
    password: '',
    confirmPassword: '',
    remmember: false,
};
