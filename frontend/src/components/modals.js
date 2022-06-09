import { Form } from "./form"
import { useEffect, useState } from "react";

export const Modal = ({title, subtitle, classes, fields, content, btnTitle, image, imageMobile, opened, close}) => {
  return (<div className={opened?`modal-full ${classes} open`:`modal-full ${classes}`}>
          <div className="modal-full__close js-close-modal" onClick={()=>close()}>
            <div className="close-icon" >
              <svg className="svg"><use xlinkHref="img/sprite.svg#close"></use></svg>
            </div>
          </div>
          <div className={image?"modal-centered modal-centered--form modal-centered--bg":"modal-centered modal-centered--form"}>
            <div className="modal-centered-wrapper">
              <div className="modal-header">
                <div className="modal-header__title h1">{title}</div>
                {subtitle?<div className="modal-text-right">{subtitle}</div>:""}
                <div className="modal-image"><img src={imageMobile} alt="" /></div>
              </div>
              {fields.length>0?<div className="modal-content">
                <Form fields={fields} btnTitle={btnTitle} />
              </div>:""}
              {content?<div className="modal-content">
                <div className="modal-text-body">
                  {content}
                </div>
              </div>:""}
            </div>
          </div>
          <div className="modal-bg"><img src={image} alt="" /></div>
        </div>
    )
}