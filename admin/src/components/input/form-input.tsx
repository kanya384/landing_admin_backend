export interface InputProps {
  title: string,
  type: string,
  placeholder: string,
  name: string,
  value: any, 
  isError: boolean,
  stateAction?: React.Dispatch<any>,
  ref?: React.MutableRefObject<any>,
}

export const FormInput: React.FC<InputProps> = (props:InputProps) => {
  let {title, name, type, placeholder, isError, value} = props
  const updateState = (name:string, value: any) => {
    props.stateAction!((prevValue: any)=>{
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
    console.log(e.target.value)
    props.stateAction!((prevValue: any)=>{
      return {
        ...prevValue,
        [name]: e.target.value === "1"
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
    default:
      return  <div className="mb-3">
                <label className="form-label" htmlFor={name}>{title}</label>
                <input className="form-control" type={type} id={name} name={name} placeholder={placeholder} value={props.value} onChange={(e)=>updateState(name, e.target.value)}/>
                {isError?"":""}
              </div> 
      
  }

}
