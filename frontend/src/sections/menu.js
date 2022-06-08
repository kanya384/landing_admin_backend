export const Menu = () => {
  return(
    <header class="header">
      <div class="wrapper">
        <div class="logo">
          <a href="#">
            <img src="./img/logo.svg" alt="logo" />
          </a>
        </div>
        <div class="header-phone">
          <div class="phone">+7 987 654-32-10</div>
          <a href="#" class="btn-recall">Заказать звонок</a>
        </div>
        
        <a href="javascript:void(0);" id="hamburger-icon">
          <span class="line line-1"></span>
          <span class="line line-2"></span>
          <span class="line line-3"></span>
          <strong>МЕНЮ</strong>
        </a>
      <nav class="navigation">
        <ul>
          <li><a href="">Инфраструктура</a></li>
          <li><a href="">Планировки и цены</a></li>
          <li><a href="">Способы покупки</a></li>
          <li><a href="">О застройщике</a></li>
          <li><a href="">Контакты</a></li>
        </ul>
      </nav>
    </div>
  </header>
  )
}