
export default function BranchIcon(props: {
    color?: string;
    width?: number;
    height?: number;
    extra?: string;
}) {
    const { color, extra, width, height, ...rest } = props;
    return (
        <svg id="elLIF4QbO201" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 125 85" shapeRendering="geometricPrecision" textRendering="geometricPrecision" width={width} height={height}>
            <path d="M1.5,2C0.671573,2,0,2.671573,0,3.5v9c0,.828427.671573,1.5,1.5,1.5h.191l1.862-3.724c.084771-.16917.257778-.275994.447-.276h8c.189222.000006.362229.10683.447.276L14.31,14h.191c.828427,0,1.5-.671573,1.5-1.5v-9c0-.397998-.158172-.779681-.439693-1.061014s-.663309-.439251-1.061307-.438986h-13ZM4,7c-.552285,0-1-.447715-1-1s.447715-1,1-1s1,.447715,1,1-.447715,1-1,1Zm8,0c-.552285,0-1-.447715-1-1s.447715-1,1-1s1,.447715,1,1-.447715,1-1,1ZM6,6c0-.552285.447715-1,1-1h2c.552285,0,1,.447715,1,1s-.447715,1-1,1h-2c-.552285,0-1-.447715-1-1Z" transform="matrix(7.812011 0 0 7.001193 0-13.736002)" fill={color} strokeMiterlimit="2" />
            <path d="M13.191,14l-1.5-3h-7.382l-1.5,3h10.382Z" transform="matrix(7.710742 0 0 6.94982 0.814058-15.354196)" fill={color}/>
            <rect width="11.305121" height="0.458254" rx="0" ry="0" transform="matrix(7.081121 0 0 9.34133 22.473536 79.999998)" fill={color} strokeWidth="0" />
        </svg>
    );
}

