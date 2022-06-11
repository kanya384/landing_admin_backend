import { useSendForm } from "../hooks/send-from.hook"
import { PhoneInput } from "./phone-input"

export const Form = ({fields, description, btnTitle, close}) => {
    const sendForm = useSendForm()
    return ( <form className="form-ec">
                <div className="form-ec__content">
                    {fields.map(({name, title, placeholder}) => {
                        if (name==="phone") {
                            return <PhoneInput />
                        }
                        return <div className="form-ec__input-row">  
                                <div className="inp-group">
                                    <div className="inp-group-label">{title}</div>
                                    <input className="input" placeholder={placeholder} name={name} />
                                </div>
                            </div>
                    })}
                    <input type="hidden" className="text" value={description} />
                    <div className="form-ec__b-btn">
                        <button className="btn-submit" onClick={(e) => sendForm.sendForm(e, close)}>
                            <span className="btn-submit-text">{btnTitle}</span>
                        </button>
                    </div>
                </div>
            </form>
    )
}