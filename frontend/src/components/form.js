import { PhoneInput } from "./phone-input"

export const Form = ({fields, btnTitle}) => {
    return ( <div class="form-ec">
                <div class="form-ec__content">
                    {fields.map(({name, title, placeholder}) => {
                        if (name==="phone") {
                            return <PhoneInput />
                        }
                        return <div class="form-ec__input-row">  
                                <div class="inp-group">
                                    <div class="inp-group-label">{title}</div>
                                    <input class="input" placeholder={placeholder} name={name} />
                                </div>
                            </div>
                    })}
                    <div class="form-ec__b-btn">
                        <button class="btn-submit">
                            <span class="btn-submit-text">{btnTitle}</span>
                        </button>
                    </div>
                </div>
            </div>
    )
}