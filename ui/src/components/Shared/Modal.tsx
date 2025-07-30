'use client'

import { useState } from 'react'
import { Dialog } from '@headlessui/react'

function Modal(props: {
    openState: boolean,
    description?: string,
    title?: string,
    children?: JSX.Element | any[];
    [x: string]: any;
}) {
    const { openState,description,title, children } = props;

    let [isOpen, setIsOpen] = useState(openState)

    console.log('open state')
    return (
        <Dialog open={isOpen} onClose={() => setIsOpen(false)}>
            <Dialog.Panel>
                <Dialog.Title>{title}</Dialog.Title>
                <Dialog.Description>
                    {description}
                </Dialog.Description>
                {children}
            </Dialog.Panel>
        </Dialog>
    );
}
export default Modal;
