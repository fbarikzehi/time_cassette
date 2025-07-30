
import toast from 'react-hot-toast';
import { DispatchRequest, RequestMethods } from "../utils/api.utils";
import { ApiUrls } from "../utils/api.url.utils";

export const AuthService = () => {

    const signup = async (payload) => {
        const toastId = toast.loading('Server is creating your workstation');
        const response = await DispatchRequest(RequestMethods.Post, ApiUrls.Auth.Signup, payload, false)
        console.log(response);
        response.meta.result ? toast.success(response.meta.messages, {
            id: toastId,
        }) : toast.error(response.meta.messages, {
            id: toastId,
        });
        return response.meta.result;
    }
    const login = async (payload) => {
        const toastId = toast.loading('Server is preparing your workstation');
        const response = await DispatchRequest(RequestMethods.Post, ApiUrls.Auth.Login, payload, false)

        response.meta.result ? (() => {
            toast.success(response.meta.messages, {
                id: toastId,
            })
        }) : toast.error(response.meta.messages, {
            id: toastId,
        });

        return response;

    }

    return {
        signup,
        login
    }
}