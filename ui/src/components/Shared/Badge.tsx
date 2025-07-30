


export default function Badge(props: {
    color?: string;
    variant?: string;
    text?: string;
    children?: JSX.Element | undefined;
    extra?: string;
}) {
    const { color,variant,text, children, extra, ...rest } = props;
    return (
        <div className={`bg-blue-600 flex flex-row items-center uppercase dark:bg-brand-400 rounded-lg py-2 px-3 text-xs font-bold text-white transition duration-200 ${extra}`} {...rest}>
            {children ? (
                <div className="flex h-5 w-5 items-center justify-center">
                    {children}
                </div>
            ) : ('')}
            {text}
        </div>
    );
}