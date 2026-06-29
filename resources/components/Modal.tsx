import { forwardRef, useImperativeHandle, useRef, useState } from "react";
import { useClickOutside } from "../hooks/useClickOutside";

type Props = {
  open: boolean;
  onClose: () => void;
  children: React.ReactNode;
};

export type ModalRef = {
    close: () => void
}

const Modal = forwardRef<ModalRef, Props>(({open, onClose, children}, ref) => {
    const [anim, setAnim] = useState<boolean>(false)
    const modalref = useRef(null)
    
    function close() {
        setAnim(true)
        
        setTimeout(() => {
            onClose()
            setAnim(false)
        }, 400);
        
    }

    useImperativeHandle(ref, () => ({
        close
    }))

    useClickOutside(modalref, () => close())
    
    if (!open) return null;

    return (
        <div className={`modal-bg`}>
            <div ref={modalref} className={`card animation-open-modal ${anim && 'animation-close-modal'}`}>
                {children}
            </div>
        </div>
    );
})

export default Modal