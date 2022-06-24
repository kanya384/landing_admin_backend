import React, { useState } from "react"
import { EditableText } from "../components/editable-text"
import { Modal } from "../components/modals"
import { Form } from "../components/form";

export const Question = () => {
  const [isOpen, setIsOpen] = useState(false)
  return ( <React.Fragment><div className="lvl15">
      <div className="wrapper">
        <div className="form-ec">
          <div className="form-ec__title"><EditableText id={"62aef61ba26e626025a8d8dd"} defaultText={"Если у вас остались вопросы, наши менеджеры вам помогут"}/></div>
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
                btnTitle={"Задать вопрос"} 
                description={`Задать вопрос менеджеру`}
                celtype={"getQuestion"}
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

export default Question