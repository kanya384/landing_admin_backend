import { EditableText } from "../components/editable-text"
import { YMaps, Map, ZoomControl, Placemark } from 'react-yandex-maps';

export const Infra = () => {
  const icons = [
    {
      location: [44.860734, 37.36766],
      hint: "МБОУ Средняя школа № 11",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.886739, 37.308678],
      hint: "МБОУ Средняя общеобразовательная школа № 1 имени Н. М. Самбурова",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.887110, 37.308170],
      hint: "МБОУ СОШ № 1",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.889328, 37.313401],
      hint: "МБОУ Гимназия Аврора имени И. И. Ладутько",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.893422, 37.316724],
      hint: "МБОУ Средняя общеобразовательная школа № 2 имени в. Каширина",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.901369, 37.332372],
      hint: "МБОУ Средняя общеобразовательная школа № 5 имени К. Соловьяновой",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.884659, 37.333254],
      hint: "МАОУ Средняя общеобразовательная школа № 6 имени Д. С. Калинина",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.888691, 37.330460],
      hint: "МБОУ Средняя общеобразовательная школа № 4 имени Героя РФ В. М. Евскина",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.883690, 37.324447],
      hint: "Школа № 13",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },

    {
      location: [44.877206, 37.330550],
      hint: "МБОУ Гимназия Эврика имени В. А. Сухомлинского",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.876892, 37.330567],
      hint: "МБОУ Средняя общеобразовательная школа № 7 имени Л. И. Севрюкова",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },
    {
      location: [44.877353, 37.316593],
      hint: "НЧОУ СОШ Светоч с углубленным изучением иностранных языков имени А. Доронина",
      icon: 'img/icon-map-school.svg',
      type: "school",
    },


    {
      location: [44.875512, 37.313726],
      hint: "Детский сад",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.870330, 37.322635],
      hint: "Карусель",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.878450, 37.322677],
      hint: "Совёнок",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.877508, 37.327519],
      hint: "Wellness Mama",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.878112, 37.329226],
      hint: "МАДОУ детский сад № 18 Виктория",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.879896, 37.323690],
      hint: "Частный детский сад Карусель ",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.880435, 37.326546],
      hint: "Юла",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.880518, 37.321058],
      hint: "Я расту",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.876210, 37.333615],
      hint: "Центр-студия Сёмушка",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.874975, 37.340629],
      hint: "МАДОУ Волшебная Страна",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.881571, 37.334553],
      hint: "МБДОУ детский сад № 3 Звездочка",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.883855, 37.331790],
      hint: "Детский сад № 12 Солнышко",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.887186, 37.320944],
      hint: "МБДОУ детский сад № 13 Теремок",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.891324, 37.326928],
      hint: "Василёк",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.891707, 37.308052],
      hint: "МБДОУ Детский сад № 17 Колобок",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.892433, 37.304242],
      hint: "МАДОУ детский сад № 6 Ракета",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.891149, 37.300622],
      hint: "МБДОУ детский сад № 16 Пчёлка",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.883532, 37.306787],
      hint: "МБДОУ детский сад № 5 Волна",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.881570, 37.315003],
      hint: "МБДОУ детский сад № 10 Светлячок",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.883268, 37.334609],
      hint: "МБДОУ детский сад № 14 Тополек",
      icon: 'img/icon-map-kids.svg',
      type: "childhood",
    },
    {
      location: [44.875499, 37.317025],
      hint: "Супермаркет Перекрёсток",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.875996, 37.325685],
      hint: "Супермаркет Пятёрочка",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.878504, 37.322376],
      hint: "Агрокомплекс",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.876115, 37.324990],
      hint: "Супермаркет Магнит",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.875026, 37.329301],
      hint: "Молочный магазин, магазин мяса Ферма",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.876866, 37.329384],
      hint: "АнапаФрудТорг",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.877610, 37.326338],
      hint: "Свежее мясо",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.878257, 37.327312],
      hint: "Мини Маркет",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.880942, 37.320738],
      hint: "Супермаркет Пятёрочка",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.880619, 37.325550],
      hint: "Супермаркет  Магнит",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.881394, 37.326906],
      hint: "Мир индейки",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    
    {
      location: [44.881590, 37.329189],
      hint: "ОптовичОк",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.879562, 37.319346],
      hint: "Свежее мясо",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.879815, 37.324136],
      hint: "Продукты",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.882008, 37.330109],
      hint: "Дары морей",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.877695, 37.334900],
      hint: "Супермаркет ВкусВилл",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.880250, 37.339849],
      hint: "Продуктовый гипермаркет Магнит экстра",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.881166, 37.336850],
      hint: "Фасоль",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.880051, 37.334283],
      hint: "Аквакультура-Кубань Океан",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.877481, 37.339314],
      hint: "Сыры и колбасы",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.876653, 37.341907],
      hint: "Супермаркет Магнит",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.876324, 37.342483],
      hint: "Океан Камчатка Сахалин",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.874831, 37.335261],
      hint: "Супермаркет Пятёрочка",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.870895, 37.335006],
      hint: "Супермаркет Пятёрочка",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.871823, 37.329061],
      hint: "Рыбное Место",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.872465, 37.328489],
      hint: "Минимаркет",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.871391, 37.329467],
      hint: "Магазин мяса, колбас Благо",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },
    {
      location: [44.877588, 37.316313],
      hint: "Супермаркет  Магнит",
      icon: 'img/icon-map-market.svg',
      type: "market",
    },

    {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Фитофарм", 
      location: [44.875059, 37.317276],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Альфа Фарма", 
      location: [44.877496, 37.316682],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Пульс", 
      location: [44.878480, 37.321404],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Магнит Аптека", 
      location: [44.876081, 37.324983],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Фитофарм", 
      location: [44.875561, 37.326049],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Наутилус", 
      location: [44.872852, 37.329726],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Здесь аптека", 
      location: [44.871687, 37.329247],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Здоровье",
      location: [44.875390, 37.329453],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Апрель", 
      location: [44.877884, 37.334706],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Будь Здоров!", 
      location: [44.878766, 37.336650],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Магнит Аптека", 
      location: [44.876061, 37.325029],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Азбука здоровья",
      location: [44.880543, 37.323287],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Советская аптека", 
      location: [44.882383, 37.322397],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Флора-Фарм", 
      location: [44.883154, 37.325679],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Фитофарм", 
      location: [44.881189, 37.326488],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Дешёвая аптека", 
      location: [44.880765, 37.325871],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Фитофарм", 
      location: [44.881997, 37.328388],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Альфа", 
      location: [44.880149, 37.329879],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Альфа", 
      location: [44.881902, 37.330500],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Аптека Здоровье", 
      location: [44.882914, 37.329424],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Альфа", 
      location: [44.882659, 37.332235],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Апрель", 
      location: [44.877872, 37.334689],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Будь Здоров!", 
      location: [44.878751, 37.336583],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Миракс-Фарм", 
      location: [44.884106, 37.326188],
  },
  {
      icon: 'img/icon-map-apteka.svg',
      type: "apteka",
      hint: "Магнит Аптека", 
      location: [44.883944, 37.320146],
  },
  {
    icon:'img/icon-map-park.svg',
    type: 'park',
    hint:"Ореховая роща",
    location:[44.876845, 37.313748],
  },
  {
    icon:'img/icon-map-park.svg',
    type: 'park',
    hint:"сквер имени В.Н. Аванесова",
    location:[44.882514, 37.305658],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
  hint:"Сквер Гудовича",
  location:[44.888081, 37.301252],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
  hint:"сквер имени П.С. Саркисьяна",
  location:[44.883397, 37.326198],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
  hint:"сквер Строитель",
  location:[44.883487, 37.332679],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
  hint:"Императорский парк",
  location:[44.873748, 37.341713],

  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
      hint:"сквер Воинской Славы", 
      location: [44.857724, 37.362855],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
      hint:"Парк ",
      location:[44.894923, 37.316183],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
      hint:"Сквер Аллея памяти",
      location:[44.895219, 37.304322],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
      hint:"парк Курортный",
      location:[44.895346, 37.299613],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
      hint:"парк 30 лет Победы",
      location:[44.896286, 37.313624],
  },
  {
      icon:'img/icon-map-park.svg',
      type: 'park',
      hint:"парк Морской порт",
      location:[44.897137, 37.305613],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Nacional", 
    location: [44.875973, 37.312125],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Бар, паб Кусто", 
    location: [44.875725, 37.312514],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Море Плова", 
    location: [44.877389, 37.317346],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Караоке-клуб, ресторан Scarlet", 
    location: [44.875741, 37.325946],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Устрица под сЫром", 
    location: [44.878541, 37.322398],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Timer", 
    location: [44.880365, 37.324816],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Золотой дракон", 
    location: [44.878246, 37.335346],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан, кальян-бар Блек Реббит", 
    location: [44.880363, 37.338931],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Subway", 
    location: [44.880128, 37.339284],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Бургер Кинг", 
    location: [44.880009, 37.339478],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"KFC Авто", 
    location: [44.879721, 37.338124],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint: "Biblioteka", 
    location: [44.883857, 37.338444],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Кафе The Room", 
    location: [44.886378, 37.325315],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Shaurma Palace", 
    location: [44.888404, 37.337426],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Cheezetta", 
    location: [44.892160, 37.329034],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Касса ", 
    location: [44.897615, 37.324984],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Кафе Галерея", 
    location: [44.895160, 37.321785],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Бар, паб Yummy Oysters", 
    location: [44.896632, 37.319975],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Хинкальная", 
    location: [44.897125, 37.318831],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Кафе Бумеранг", 
    location: [44.897274, 37.318069],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Облака", 
    location: [44.897935, 37.319512],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Oscar", 
    location: [44.897462, 37.319544],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Корона", 
    location: [44.897338, 37.319316],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Золотой пляж", 
    location: [44.899946, 37.316238],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Green Duck", 
    location: [44.900620, 37.321138],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Японsky", 
    location: [44.898408, 37.333354],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Кафе Али-Баба", 
    location: [44.901473, 37.336854],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Beershop ДПЗ", 
    location: [44.900935, 37.334850],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Shamsi Hall", 
    location: [44.896631, 37.340827],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Макдоналдс", 
    location: [44.899873, 37.348151],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Бар, паб Пивная демократия", 
    location: [44.873069, 37.329434],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Кафе Старый Хутор", 
    location: [44.877471, 37.338650],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан-бар Biblioteka", 
    location: [44.883818, 37.338404],
  },
  {
    icon:'img/icon-map-cafe.svg',
    type:"cafe",
    hint:"Ресторан Мацони", 
    location: [44.896142, 37.306579],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
        hint:"Sunfit", 
        location: [44.875997, 37.325515],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Спортклуб Самурай Анапа", 
    location: [44.876177, 37.324927],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Подиум", 
    location: [44.877811, 37.322131],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Fresh", 
    location: [44.877502, 37.325290],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Fitness SPA kenguru", 
    location: [44.877080, 37.333181],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Фитнес-студия Fresh Style Фреш Стайл г. Анапа", 
    location: [44.879959, 37.325409],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Ось мира", 
    location: [44.880876, 37.330474],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"FitZone", 
    location: [44.880106, 37.336046],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Arny-Gym ", 
    location: [44.881384, 37.323657],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Стройность", 
    location: [44.887149, 37.329598],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Стиль-Класс", 
    location: [44.887438, 37.329940],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Кузня", 
    location: [44.888711, 37.311748],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Грация", 
    location: [44.889314, 37.310178],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Пульс 120", 
    location: [44.887792, 37.302919],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"@primetime. ", 
    location: [44.893076, 37.315285],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Тренажерный зал для женщин", 
    location: [44.893261, 37.332358],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Genetika", 
    location: [44.891044, 37.336711],
  },
  {
    icon:'img/icon-map-fitness.svg',
    type: "fitness",
    hint:"Dzen", 
    location: [44.892578, 37.339584],
  },

    {
      icon:'img/icon-map-bus.svg',
      type: "bus",
      hint:"Оставновка", 
      location: [44.878135, 37.324008],
    },
    {
      icon:'img/icon-map-bus.svg',
      type: "bus",
      hint:"Оставновка", 
      location: [44.876531, 37.325671],
    },
    {
      icon:'img/icon-map-bus.svg',
      type: "bus",
      hint:"Оставновка", 
      location: [44.871763, 37.322847],
    },
    {
      icon:'img/icon-map-beach.svg',
      type: "beach",
      hint:"Пляж", 
      location: [44.878636, 37.308242],
    },

    {
      icon:'img/icon-map-beach.svg',
      type: "beach",
      hint:"Пляж", 
      location: [44.882121, 37.305335],
    },

  ]
  return (
    <div className="lvl13">
      <div className="wrapper">
        <div className="h-title"><EditableText id={"62aef61ba26e626025a8d8dc"} defaultText={"Развитая <br />инфраструктура"}/></div>
        <div className="map-info" style={{width:"100%", height: 704, display:"block", position: "relative", overflow: "hidden" }}>
            <YMaps>
                    <Map defaultState={{ center: [44.887186, 37.320944], zoom: 15, controls: [] }} style={{ width: "100%", height: "100%", position: "absolute" }}  >
                        <ZoomControl options={{ float: 'left' }} />
                        {icons.map((icon)=>{
                          return <Placemark geometry={icon.location}
                                    options={{
                                        iconLayout: 'default#image',
                                        hideIconOnBalloonOpen: false,
                                        iconImageSize: [50, 50],
                                        iconImageOffset: [-50, -72],
                                        cursor: 'default',
                                        iconShadow: true,
                                        balloonclose: true,
                                        iconImageHref: icon.icon,
                                    }}
                                    properties= {{
                                      hintContent: icon.hint,
                                    }}
                                    modules= {
                                        ['geoObject.addon.balloon', 'geoObject.addon.hint']
                                    }
                                />
                        })}
                        
                    </Map>

            </YMaps>
        </div>
        <div className="map-info-legend">
          <div className="map-info-legend__item">
            <div className="icon"><img src="img/i-school.svg" alt="" /></div>
            <div className="title">Школы</div>
          </div>
          <div className="map-info-legend__item">
            <div className="icon"><img src="img/i-pharmacy.svg" alt="" /></div>
            <div className="title">Аптеки</div>
          </div>
          <div className="map-info-legend__item">
            <div className="icon"><img src="img/i-parks.svg" alt="" /></div>
            <div className="title">Парки</div>
          </div>
          <div className="map-info-legend__item">
            <div className="icon"><img src="img/i-kindergartens.svg" alt="" /></div>
            <div className="title">Детские сады</div>
          </div>
          <div className="map-info-legend__item">
            <div className="icon"><img src="img/i-market.svg" alt="" /></div>
            <div className="title">Магазины</div>
          </div>
          <div className="map-info-legend__item">
            <div className="icon"><img src="img/i-fitness.svg" alt="" /></div>
            <div className="title">Фитнес-центры</div>
          </div>
          <div className="map-info-legend__item">
            <div className="icon"><img src="img/i-cafe.svg" alt="" /></div>
            <div className="title">Кафе и рестораны</div>
          </div>
          <div class="map-info-legend__item">
            <div class="icon"><img src="img/i-beach.svg" alt="" /></div>
            <div class="title">Пляжи</div>
          </div>
          <div class="map-info-legend__item">
            <div class="icon"><img src="img/i-bus.svg" alt="" /></div>
            <div class="title">Остановки</div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Infra