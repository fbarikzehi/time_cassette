


export function classNames(...classes: string[]) {
    return classes.filter(Boolean).join(' ')
}

export enum positions {
    Right,
    Left,
    Top,
    Bottom,
    Center,
    TopRight,
    TopLeft,
    BottomRight,
    BottomLeft,
}