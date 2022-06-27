import React from "react"
import { useSendForm } from "../hooks/send-from.hook"
import { PhoneInput } from "./phone-input"
import { RangeSliderPopup } from "./range-slider-popup"

export const Form = ({fields, description, btnTitle, celtype, close, callback}) => {
    const sendForm = useSendForm()
    if (celtype=="getPodbor" || celtype=="getExcursion" || celtype === "getQuestionAbout"|| celtype == "getFlat" || celtype == "getPresentation" || celtype == "getQuestion" ) {
        return  <React.Fragment>
                    <form className={celtype == "getFlat"?"room-card__form":"form-ec__content"}>
                        <div className="form-ec__input-row">
                            {fields.map(({name, title, type, placeholder}) => {
                                if (name==="phone") {
                                    return <PhoneInput alt={true} />
                                }
                                
                                return <div className="inp-group">
                                        <div className="inp-group-label">{title}</div>
                                        <input className="input" placeholder={placeholder} name={name} type="text" />
                                    </div>
                            })}
                            <input type="hidden" className="text" value={description} />
                            {celtype=="getPodbor"?<div className={celtype=="getPodbor"?"inp-group":"form-ec__b-btn"}>
                                <button className="btn-submit" celtype={celtype} onClick={(e) => sendForm.sendForm(e, callback)}><span className="btn-submit-text">{btnTitle}</span></button>
                            </div>:""}
                        </div>
                        {celtype!=="getPodbor"?<div className="form-ec__b-btn"><button className="btn-submit" celtype={celtype} onClick={(e) => sendForm.sendForm(e, callback)}><span className="btn-submit-text">{btnTitle}</span></button></div>:""}
                    </form>

                </React.Fragment>

    }
    return ( <form className={"form-ec"}>
                <div className="form-ec__content">
                    {fields.map(({name, title, type, placeholder, min, max, step, postfix, value, callback}) => {
                        if (name==="phone") {
                            return <PhoneInput />
                        }
                        if (type === "time") {
                            return <div className="form-ec__input-row">
                                        <div className="inp-group input-group-x2__col">
                                            <div className="inp-group-label">Часы</div>
                                            <input className="input dop-info-time" data={"Час"} placeholder="18 ч" />
                                        </div>
                                        <div className="inp-group input-group-x2__col">
                                            <div className="inp-group-label">Минуты</div>
                                            <input className="input dop-info-time" data={":"} placeholder="00 мин" />
                                        </div>
                                    </div>
                        }
                        if (type === "slider") {
                            return <div className="form-ec__input-row">
                                        <div className="inp-group">
                                            <div className="inp-group__title">{title}</div>
                                            <RangeSliderPopup
                                                value={value}
                                                min={min}
                                                max={max}
                                                step={step}
                                                callback={callback}
                                                postfix={postfix}
                                            />
                                        </div>
                                    </div>
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
                        <button className="btn-submit" celtype={celtype} onClick={(e) => sendForm.sendForm(e, close)}>
                            <span className="btn-submit-text">{btnTitle}</span>
                        </button>
                    </div>
                </div>
            </form>
    )
}
