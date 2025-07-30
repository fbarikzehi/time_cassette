

const formatDomain = (name) => {
    // return `https://api.${name}.timecassette.com`
    return `https://api.timecassette.com`
}

const domains = {
    default: formatDomain(),
    auth: formatDomain("auth"),
    cassette: formatDomain("cassette"),
    fragment: formatDomain("fragment"),
    branch: formatDomain("branch"),
    time: formatDomain("time"),
}

/**
 * Ex:ApiUrls.Auth.Login
 */
export const ApiUrls = {
    Auth: {
        Login: `${domains.default}/login`,
        Logout: `${domains.default}/logout`,
        Signup: `${domains.default}/signup`,
    },
    User: {
        SearchEmail: `${domains.default}/user/searchemail`,
    },
    Cassette: {
        Cassettes: `${domains.default}/cassettes`,
        Create: `${domains.default}/cassettes/create`,
        Update: `${domains.default}/cassettes/cassette/update`,
        Delete: `${domains.default}/cassettes/cassette/delete`,
        Fragments: `${domains.default}/cassettes/fragments`,
    },
    Fragment: {
        Create: `${domains.default}/cassettes/fragments/create`,
        Update: `${domains.default}/cassettes/fragments/fragment/update`,
        Delete: `${domains.default}/cassettes/fragments/fragment/delete`,
        Branches: `${domains.default}/cassettes/fragments/branches`,
    },
    Branch: {
        Create: `${domains.default}/cassettes/fragments/branches/create`,
        Update: `${domains.default}/cassettes/fragments/branches/branch/update`,
        DeleteRequest: `${domains.default}/cassettes/fragments/branches/branch/delete/request`,
        DeleteConfirm: `${domains.default}/cassettes/fragments/branches/branch/delete/confirm`,
        Times: `${domains.default}/cassettes/fragments/branches/times`,
    },
    Time: {
        Create: `${domains.default}/cassettes/fragments/branches/times/create`,
        Update: `${domains.default}/cassettes/fragments/branches/times/time/update`,
        Delete: `${domains.default}/cassettes/fragments/branches/times/time/delete`,
        DeleteAll: `${domains.default}/cassettes/fragments/branches/times/delete`,
        UpdateStart: `${domains.default}/cassettes/fragments/branches/times/time/update/start`,
        UpdateEnd: `${domains.default}/cassettes/fragments/branches/times/time/update/end`,
        UpdateDescription: `${domains.default}/cassettes/fragments/branches/times/time/update/description`,
    }
}
