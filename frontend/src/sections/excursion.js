export const Excursion = () => {
  return ( <div class="lvl5">
      <div class="wrapper">
        <div class="form-ec">
          <div class="form-ec__title">Экскурсия в жилой комплекс</div>
          <div class="form-ec__content">
            <div class="form-ec__input-row">
              <div class="inp-group">
                <div class="inp-group-label">Ваше имя</div>
                <input class="input" placeholder="Ваше имя" type="text" />
              </div>
              <div class="inp-group">
                <div class="inp-group-label">Телефон</div>
                <input class="input" placeholder="Телефон" type="text" />
              </div>
            </div>
            <div class="form-ec__b-btn"><button class="btn-submit"><span class="btn-submit-text">Отправить</span></button></div>
          </div>
        </div>
      </div>
    </div>
  )
}