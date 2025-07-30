

const QueryParamsMapper = (params: any) => {
    return params ? Object.keys(params).map((key: string) => `${encodeURIComponent(key)}=${encodeURIComponent(params[key])}`).join('&') : '';
}
/**
 * GET, POST, PUT ,PATCH , DELETE
 */
const enum RequestMethods {
    Get = "GET",
    Post = "POST",
    Put = "PUT",
    Patch = "PATCH",
    Delete = "DELETE",
};

const Headers = (auth: boolean) => {
    const session = localStorage.getItem('_session');
    let headers = {
        'Content-Type': 'application/json',
        Authorization: ''
    }
    if (auth) {
        headers.Authorization = `Bearer ${session ? JSON.parse(session).token : ''}`;
    }

    return headers;
}

async function handleResponse(response: Response) {
    // console.log(response.json());
    const isJson = response.headers?.get('content-type')?.includes('application/json');
    const data = isJson ? response.json() : null;
    // console.log(data);
    if (!data)
        return {
            meta: {
                messages: ['Oops!something goes wrong with server response'],
                result: false
            },
        }

    return data;
}

async function translateError(error: any) {
    if (!error.ok) {
        if (error.status == 401) {
            //logout();
        }
        else if (error.status == 400) {
            // console.log('405');
        }
        else if (error.status == 503) {
            // console.log('405');
        }
    }

    return error;
}

/**
 * @param {TYPE} method - RequestMethods: GET, POST, PUT, DELETE
 * @param {string} url  - ApiUrls
 * @param {Object} payload - {}
 * @param {boolean} auth - true/false
 */
const DispatchRequest = async (method: RequestMethods, url: string, payload: any, auth: boolean = true) => {
    let requestInit = {
        method,
        headers: Headers(auth)
    };
    switch (method) {
        case RequestMethods.Post:
        case RequestMethods.Delete:
        case RequestMethods.Put:
            if (!payload) break;
            requestInit = { ...requestInit, ...{ body: JSON.stringify(payload) } }
            break;
        case RequestMethods.Get:
            url += (payload ? `?${QueryParamsMapper(payload)}` : '');
            break;
        default:
            break;
    }
    // console.log(requestInit)
    const response = await fetch(url, requestInit).then(handleResponse).catch((error) => {
        console.log(error)
        return {
            meta: {
                messages: [error.message],
                result: false
            },
        }
        // console.error(error)
        // const translatedError = translateError(error);
        // return Promise.resolve(translatedError);
    });
    // console.log(response)
    return response;
}


export { RequestMethods, DispatchRequest }