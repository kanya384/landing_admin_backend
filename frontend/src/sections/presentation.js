import React, { useState } from "react"
import { EditableText } from "../components/editable-text"
import { Modal } from "../components/modals"
import { Form } from "../components/form";

export const Presentation = () => {
  const [isOpen, setIsOpen] = useState(false)
  return ( <React.Fragment>
    <div className="lvl7">
      <div className="wrapper">
        <div className="form-ec">
          <div className="form-ec__title"><EditableText id={"62aef61ba26e626025a8d8cb"} defaultText={"Получите подробную презентацию жилого комплекса"}/></div>
          <div className="form-ec__content">
              <Form fields={[
                    {
                      type:"text",
                      name: "name",
                      placeholder: "Имя",
                      required: false,
                    },
                    {
                        type:"text",
                        name: "phone",
                        placeholder: "Телефон",
                        required: true,
                    },
                  ]} 
                  btnTitle={"Получить презентацию"} 
                  description={`Получить подробную презентацию жилого комплекса`}
                  celtype={"getPresentation"}
                  close={()=>{}} 
                  callback={()=>{setIsOpen(true)}}
              />

          </div>
        </div>
      </div>
    </div>
    <Modal 
          success={true}
          position={window.pageYOffset}
          opened={isOpen}
          close = {()=>{setIsOpen(null)}}
        />
    </React.Fragment>
  )
}

export default Presentation