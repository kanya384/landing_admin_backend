import React from "react";
import { FormInput, InputProps } from "../input";
import { Popup } from "../popup/popup";
interface PopupProps {
  Forms: Form[],
  Opened: number| null,
  Toggle: any,
}
interface Form {
  title: string,
  error: string,
  stateAction: React.Dispatch<any>,
  sendAction: ()=>void,
  fields: InputProps[],
}
export const Popups: React.FC<PopupProps> = (props: PopupProps) => {
  return(
    <React.Fragment>
    {
    props.Forms.map((form, index) => {
      return <Popup title={form.title} key={index} isOpen={index===props.Opened} toggle={props.Toggle} send={form.sendAction} error={"ошибка"}>
        {form.error !== ""?<div className="alert alert-soft-danger" role="alert">{form.error}</div>:<div></div>}
        {form.fields.map((field, index)=>{
          return <FormInput key={index} title={field.title} type={field.type} placeholder={field.placeholder} name={field.name} isError={field.isError} value={field.value} ref={field.ref} stateAction={form.stateAction} />
        })}
      </Popup>
    })
    }
    </React.Fragment>
  )
}