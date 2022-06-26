import React from "react"
import Dropzone from "react-dropzone"
import { Cloud, Upload } from "react-feather"

export interface InputProps {
  title: string,
  type: string,
  placeholder: string,
  name: string,
  value: any, 
  isError: boolean,
  stateAction?: React.Dispatch<any>,
  ref?: React.MutableRefObject<any>,
  fields?:{value:number | string, name:string, type:number}[],
}

export const FormInput: React.FC<InputProps> = (props:InputProps) => {
  let {title, name, type, placeholder, isError, value} = props
  const updateState = (name:string, value: any) => {
    props.stateAction!((prevValue: any)=>{
      if (type == "number") {
        return {
          ...prevValue,
          [name]: parseInt(value)
        }
      }
      return {
        ...prevValue,
        [name]: value
      }
    })
  }

  const updateStateFile = (e: any) => {
    props.stateAction!((prevValue: any)=>{
      return {
        ...prevValue,
        [name]: e.target.files[0]
      }
    })
  }

  const updateStateBoolean = (e: any) => {
    props.stateAction!((prevValue: any)=>{
      return {
        ...prevValue,
        [name]: e.target.value === "1"
      }
    })
  }
  const updateStateSelect = (e: any) => {
    props.stateAction!((prevValue: any)=>{
      if (props.fields![0].type === 0) {
        return {
          ...prevValue,
          [name]: parseInt(e.target.value)
        }
      }
      return {
        ...prevValue,
        [name]: e.target.value
      }
    })
  }

  const photosDropZoneChange = (acceptedFiles: any) => {
    props.stateAction!((prevValue: any)=>{
      console.log(acceptedFiles)
      let newArray: any[] = []
      if (prevValue[name]) {
        newArray = [...prevValue[name]]
      } 
      acceptedFiles.map((file: File) => {
        if (file.type == "image/jpeg" || file.type == "image/png") {
          newArray.push(file)
        }
      })
      return {
        ...prevValue,
        [name]: newArray
      }
    })
  }
  
  switch (type) {
    case "textarea":
      return <div className="mb-3">
                <label className="form-label" htmlFor={name}>{title}</label>
                <textarea className="form-control" id={name} name={name} rows={3} placeholder={placeholder} value={props.value} onChange={(e)=>updateState(name, e.target.value)}></textarea>
                {isError?"":""}
              </div> 
    case "file":
      return  <div className="mb-3">
                <label className="form-label" htmlFor={name}>{title}</label>
                <input className="form-control" type={type} id={name} name={name} placeholder={placeholder} ref={props.ref} onChange={updateStateFile}/>
                {isError?"":""}
              </div>
    case "boolean":
      return  <div className="mb-3">
                <label className="form-label" htmlFor={name}>{title}</label>
                <select className="form-select" onChange={updateStateBoolean} defaultValue={value===true?"1":"0"} id={name}>
                    <option value={"1"}>Да</option>
                    <option value={"0"}>Нет</option>
                </select>
                {isError?"":""}
              </div>
    case "select":
      return  <div className="mb-3">
                <label className="form-label" htmlFor={name}>{title}</label>
                <select className="form-select" onChange={updateStateSelect} defaultValue={value} id={name}>
                  <option value="">{}</option>
                  {props.fields?.map((field)=>{
                    return <option value={field.value}>{field.name}</option>
                  })}
                </select>
                {isError?"":""}
              </div>
    case "drop":
      return <Dropzone onDrop={photosDropZoneChange}>
      {({ getRootProps, getInputProps }) => (
          <div className="dropzone dz-clickable dz-started" {...getRootProps()}>
              <div className="dz-message needsclick">
                  <input {...getInputProps()} />
                  <div>
                      <div className="alert alert-soft-success d-flex flex-column justify-content-center" role="alert">
                        <h4 className="alert-heading fw-semi-bold d-flex justify-content-center"><Upload size={30} /></h4>
                        <p>Перетащите картинки в контейнер или нажмите чтобы загрузить.</p>
                        {value.length > 0 ? <React.Fragment><hr className="bg-300" /><p className="mb-0">{value.length + " файлa(-ов) на загрузку"}</p></React.Fragment> : ""}
                      </div>
                    </div>
              </div>
          </div>
      )}
      </ Dropzone>
    default:
      return  <div className="mb-3">
                <label className="form-label" htmlFor={name}>{title}</label>
                <input className="form-control" type={type} id={name} name={name} placeholder={placeholder} value={props.value} onChange={(e)=>updateState(name, e.target.value)}/>
                {isError?"":""}
              </div> 
      
  }

}
