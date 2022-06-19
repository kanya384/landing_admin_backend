import { EditableText } from "../components/editable-text"

export const Excursion = () => {
  return ( <div className="lvl5">
      <div className="wrapper">
        <div className="form-ec">
          <div className="form-ec__title"><EditableText id={"62aef61ba26e626025a8d8c8"} defaultText={"Экскурсия в жилой комплекс"}/></div>
          <div className="form-ec__content">
            <div className="form-ec__input-row">
              <div className="inp-group">
                <div className="inp-group-label">Ваше имя</div>
                <input className="input" placeholder="Ваше имя" type="text" />
              </div>
              <div className="inp-group">
                <div className="inp-group-label">Телефон</div>
                <input className="input" placeholder="Телефон" type="text" />
              </div>
            </div>
            <div className="form-ec__b-btn"><button className="btn-submit"><span className="btn-submit-text">Отправить</span></button></div>
          </div>
        </div>
      </div>
    </div>
  )
}