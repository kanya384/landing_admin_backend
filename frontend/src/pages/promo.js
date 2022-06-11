import React from "react"
import { AppBar } from "../components/appbar"
import { Footer } from "../components/footer"

export const Promo = () => {
  return (
    <React.Fragment>
      <AppBar />
      <div class="container">
        <div class="wrapper">
          <div class="content-wrapper">
            <main class="content">
              <div class="promo-first">
                <div class="promo-first__title"><h1 class="h1">выбирай жизнь на юге россии</h1></div>
                <div class="promo-first__text">
                  Анапа находится на 45 параллели, называемой «золотой линией» или «линией жизни». На этой широте
                  наиболее благоприятные условия для жизни человека. По аналогии с человеческим организмом 45 параллель
                  - это район сердца. Ипостась «золотого сечения» и гармоничного деления в размерах земного шара.
                </div>
              </div>
              <div class="about-info-list about-info-list--promo">
                <div class="about-info-item">
                  <div class="img">
                    <div class="img-float-label">
                      <img src="img/promo1.png" alt="" />
                      <div class="b-label">ехать 20 минут</div>
                    </div>
                  </div>
                  <div class="b-text">
                    <div class="title">Романтика Абрау-Дюрсо</div>
                    <div class="text">
                      Абрау-Дюрсо - так называется озеро и известный завод шампанских вин, куда можно отправиться на
                      экскурсию и дегустацию. Находится это место в получасе езды от Анапы. Территория завода
                      представляет собой благоустроенную набережную, великолепный парк с зелеными аллеями, множеством
                      кафе и ресторанов. Здесь есть сцена, где с концертами часто выступают известные исполнители, а
                      также местные кавер-группы. Это настоящая роскошь - провести вечер в Абрау-Дюрсо, любуясь закатом
                      с бокалом шампанского!
                    </div>
                  </div>
                </div>
                <div class="about-info-item">
                  <div class="img">
                    <div class="img-float-label">
                      <img src="img/promo2.png" alt="" />
                      <div class="b-label">ехать 15 минут</div>
                    </div>
                  </div>
                  <div class="b-text">
                    <div class="title">Пляж станицы Благовещен-ской</div>
                    <div class="text">
                      Это одно из самых популярных мест отдыха под Анапой. Сюда едут со всей страны: обычные туристы,
                      любители отдохнуть дикарем, дайверы, сёрферы и кайтеры. Здесь можно принимать солнечные ванны или
                      взять катамаран и отправиться в море, где сквозь толщу воды видно рыб и медуз.
                    </div>
                  </div>
                </div>
                <div class="about-info-item">
                  <div class="img">
                    <div class="img-float-label">
                      <img src="img/promo3.png" alt="" />
                      <div class="b-label">ехать 18 минут</div>
                    </div>
                  </div>
                  <div class="b-text">
                    <div class="title">Заповедник Утриш и Кипарисовое озеро</div>
                    <div class="text">
                      Заповедник Большой Утриш привлекает огромное количество туристов своей первозданной природой.
                      Именно там растут реликтовые можжевеловые леса. А также в 14 км к югу от Анапы в одном из самых
                      загадочных мест черноморского побережья, за склонами зеленых сопок, у подножья Кавказских гор,
                      расположен парк «Долина Сукко». Главная достопримечательность - 32 кипариса, которые растут в
                      водах озера.
                    </div>
                  </div>
                </div>
                <div class="about-info-item">
                  <div class="img">
                    <div class="img-float-label">
                      <img src="img/promo4.png" alt="" />
                      <div class="b-label">ехать 15 минут</div>
                    </div>
                  </div>
                  <div class="b-text">
                    <div class="title">Атамань — путешествие в прошлое</div>
                    <div class="text">
                      Настоящая казачья станица «Атамань» широко раскинулась на берегу Таманского залива. Это не
                      музейный экспонат, где ничего нельзя трогать руками, а место, где можно окунуться в быт казаков, а
                      также узнать об истории этой земли. «Атамань» перенесет своих гостей на сотни лет в прошлое и
                      позволит почувствовать дух казачества, несгибаемой воли и бескрайней свободы
                    </div>
                  </div>
                </div>
              </div>
              <div class="form-ec">
                <div class="form-ec__header">
                  <div class="form-ec__title">захотелось жить у моря?</div>
                  <div class="form-ec__text">Заполните форму ниже, и наш менеджер свяжется с вами!</div>
                </div>
                <div class="form-ec__content">
                  <div class="form-ec__input-row">
                    <div class="inp-group">
                      <div class="inp-group-label">Ваше имя</div>
                      <input class="input" placeholder="Ваше имя" />
                    </div>
                    <div class="inp-group">
                      <div class="inp-group-label">Номер телефона</div>
                      <input class="input" placeholder="Номер телефона" />
                    </div>
                  </div>
                  <div class="form-ec__b-btn">
                    <button class="btn-submit"><span class="btn-submit-text">Отправить</span></button>
                  </div>
                </div>
              </div>
            </main>
          </div>
        </div>
      </div>
      <Footer />
    </React.Fragment>
  )
}