
export enum menuPath {
    Workstation = '/workstation',
    Cassettes = '/workstation/cassettes',
    // Fragments = '/workstation/cassettes/fragments',
    // Branches = '/workstation/cassettes/fragments/branches',
    // Times = '/workstation/cassettes/fragments/branches/times',
    Teams = '/workstation/teams',
    Calendar = '/workstation/calendar',
    Reports = '/workstation/reports'
}

export const MenuItems = [
    { name: 'Workstation', href: menuPath.Workstation },
    { name: 'Cassettes', href: menuPath.Cassettes },
    // { name: 'Fragments', href: menuPath.Fragments },
    // { name: 'Branches', href: menuPath.Branches },
    // { name: 'Times', href: menuPath.Times },
    { name: 'Teams', href: menuPath.Teams },
    { name: 'Calendar', href: menuPath.Calendar },
    { name: 'Reports', href: menuPath.Reports },
]


