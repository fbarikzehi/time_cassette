
'use client'

export default function LoadingSpinner(props: { extraClasses?: string; width: number, height: number }) {
    const { extraClasses, width, height, ...rest } = props;

    return (
        <div className="flex justify-center">
            <svg width={width} height={height} viewBox="0 0 16 16" fill="none" data-view-component="true" className={`spinner ${extraClasses}`}  {...rest}>
                <circle cx="8" cy="8" r="7" stroke="currentColor" strokeOpacity="0.25" strokeWidth="2" vectorEffect="non-scaling-stroke"></circle>
                <path d="M15 8a7.002 7.002 0 00-7-7" stroke="currentColor" strokeWidth="2" strokeLinecap="round" vectorEffect="non-scaling-stroke"></path>
            </svg>
        </div>
    );
}



