export const Plans = () => {
  return (
    <div class="lvl8">
      <div class="wrapper">
        <div class="h-title">Подбор квартиры на 3D-плане</div>
        <div class="text">Выберите дом и этаж, чтобы посмотреть планировки и узнать цены. Передвигайтесь влево-вправо, ввех и вниз</div>
        <div class="b-3d">
          <img src="img/home3d.svg" alt="" />
        </div>
        <div class="b-link-row">
          <a class="lnk-params js-open-modal-flat" href="#">
            <span class="lnk-params__ico">
            <svg class="ico" width="66" height="30" viewBox="0 0 66 30"
                  xmlns="http://www.w3.org/2000/svg">
              <path d="M64.5 14.9983L65.2033 15.7092L65.9219 14.9982L65.2032 14.2874L64.5 14.9983ZM51.0509 29.7109L65.2033 15.7092L63.7967 14.2875L49.6443 28.2891L51.0509 29.7109ZM65.2032 14.2874L51.0508 0.289034L49.6444 1.71097L63.7968 15.7093L65.2032 14.2874ZM64.5 13.9983H0.5V15.9983H64.5V13.9983Z" />
            </svg>
            </span>
            Подбор по параметрам
          </a>
        </div>
      </div>
    </div>
  )
}