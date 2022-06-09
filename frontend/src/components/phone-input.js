import { useState } from "react";

export const PhoneInput = () => {
    const [phone, setPhone] = useState("")
    const phoneChange = (e) => {
        
        let lastVal = phone;
        let value = e.target.value;
        
        setPhone(value)
        if (value.length < lastVal.length && value.length < 5) {
            setPhone('');
        }

        if (value.length === 1) {
            if (parseInt(value) === 8) {
                setPhone('+7 (');
            } else if (parseInt(value) === 9) {
                setPhone('+7 (9');

            } else if (parseInt(value) === 7) {
                setPhone('+7 (');
            } else {
                setPhone('');
            }
        }

        if (value.length === 8 && value.length > lastVal.length) {
            value = value.substring(0, value.length - 1) + ') ' + value.substring(value.length - 1, value.length);
            setPhone(value);
        }
        if (value.length === 9 && value.length < lastVal.length) {
            value = value.substring(0, value.length - 3);
            setPhone(value);
        }
        if (value.length === 13 && value.length > lastVal.length) {
            value = value.substring(0, value.length - 1) + '-' + value.substring(value.length - 1, value.length);
            setPhone(value);
        }
        if ((value.length === 13) && value.length < lastVal.length) {
            value = value.substring(0, value.length - 2);
            setPhone(value);
        }
        if (value.length === 16 && value.length > lastVal.length) {
            value = value.substring(0, value.length - 1) + '-' + value.substring(value.length - 1, value.length);
            setPhone(value);
        }
        if (value.length === 16 && value.length < lastVal.length) {
            value = value.substring(0, value.length - 2);
            setPhone(value);
        }
        if (value.length > 18) {
            e.currentTarget.parentElement.classList.add("err");
        } else {
            e.currentTarget.parentElement.classList.remove("err");
        }

        if (value.length < lastVal.length) {
            if (value.length < 5) {
                setPhone('');
            }
        }
    }
    return ( <div class="form-ec__input-row">
                <div class="inp-group">
                    <div class="inp-group-label">Телефон</div>
                    <input class="input" placeholder="Телефон" onChange={phoneChange} value={phone} />
                </div>
            </div>
    )
}
