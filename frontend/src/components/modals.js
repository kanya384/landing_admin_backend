import { useEffect, useState } from "react";
import { FloorsModal } from "../feature/floors-modal";
import { Form } from "./form"
const parse = require('html-react-parser');

export const Modal = ({title, subtitle, classes, liter, success, description, fields, content, celtype, btnTitle, image, imageMobile, opened, close}) => {
  const [showSuccess, setShowSuccess] = useState(false)
  useEffect(()=>{
    setShowSuccess(false)
  },[opened])
  if (success || showSuccess) {
    return <div className={opened?`modal-full ${classes} open`:`modal-full ${classes}`} style={{zIndex:90000}}>
              <div className="modal-full__close js-close-modal" onClick={()=>close()}>
                <div className="close-icon" >
                  <svg className="svg"><path d="M22 2L12 12M12 12L2 22M12 12L22 22M12 12L2 2" stroke-width="4"></path></svg>
                </div>
              </div>
              <div className={"modal-centered modal-centered--form modal-centered--bg"}>
                <div className="modal-centered-wrapper">
                  <div className="modal-header">
                    <div class="modal-header__title h1">Спасибо</div>
                  </div>
                  <div class="modal-content">
                    <div class="modal-text-body">
                      В ближайшее время с вами свяжется специалист.
                    </div>
                  </div>
                </div>
              </div>
            </div>
  }
  
  
  if (classes === "modal-flat") {
    return <FloorsModal {...{title, liter, classes, opened, close}}/>
  }

  
  
  return (<div className={opened?`modal-full ${classes} open`:`modal-full ${classes}`}>
          <div className="modal-full__close js-close-modal" onClick={()=>{setShowSuccess(false); close(); }}>
            <div className="close-icon" >
              <svg className="svg"><use xlinkHref="img/sprite.svg#close"></use></svg>
            </div>
          </div>
          <div className={image?"modal-centered modal-centered--form modal-centered--bg":"modal-centered modal-centered--form"}>
            <div className="modal-centered-wrapper">
              <div className="modal-header">
                <div className="modal-header__title h1">{title}</div>
                {subtitle?<div className="modal-text-right">{parse(subtitle)}</div>:""}
                {imageMobile?<div className="modal-image"><img src={imageMobile} alt="" /></div>:""}
              </div>
              {fields.length>0?<div className="modal-content">
                <Form fields={fields} btnTitle={btnTitle} description={description} celtype={celtype} close={() => {setShowSuccess(true)}} />
              </div>:""}
              {content?<div className="modal-content">
                <div className="modal-text-body">
                  {parse(content)}
                </div>
              </div>:""}
            </div>
          </div>
          <div className="modal-bg"><img src={image} alt="" /></div>
        </div>
    )
}