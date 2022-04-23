import { X } from 'react-feather';
import { Modal, ModalBody } from "reactstrap";

interface PopupProps {
    title: string,
    children: any,
    isOpen: boolean,
    send: ()=>void,
    error: string,
    toggle: any
}
export const Popup: React.FC<PopupProps> = ({title, children, isOpen, send, toggle}) => {
    return (<div>
        <Modal isOpen={isOpen} toggle={toggle}>
            <ModalBody>
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title" id="exampleModalLabel">{title}</h5>
                        <button className="btn p-1" type="button" data-bs-dismiss="modal" onClick={()=>toggle()} aria-label="Close"><X /></button>
                    </div>
                    <div className="modal-body">
                        {children}
                    </div>
                    <div className="modal-footer">
                        <button className="btn btn-primary" type="button" onClick={()=>send()}>Отправить</button>
                        <button className="btn btn-outline-primary" type="button" data-bs-dismiss="modal" onClick={()=>toggle()}>Отмена</button>
                    </div>
                </div>
            </ModalBody>
        </Modal>
    </div>)
}