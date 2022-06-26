import { Layers, Tool, Layout, Youtube, Grid, BookOpen, Server, MessageCircle, RefreshCw, Settings, LogOut } from 'react-feather';

interface MenuItem {
  name: string;
  icon: JSX.Element | null;
  url: string | null;
  children: MenuItem[] | null;
  divider: boolean ,
}

const ICON_SIZE = 18

const elements: MenuItem[] = [
  {
    name: "Аналитика",
    icon: null,
    url: null,
    children: null,
    divider: true,
  },
  {
    name: 'Панель управления',
    icon:  <Grid size={ICON_SIZE} />,
    url: "/",
    children: [],
    divider: false
  },
  {
    name: "Контент",
    icon: null,
    url: null,
    children: null,
    divider: true,
  },
  {
    name: 'Постеры',
    icon:  <Layers size={ICON_SIZE} />,
    url: "/posters",
    children: [],
    divider: false
  },
  {
    name: 'О проекте',
    icon:  <MessageCircle size={ICON_SIZE} />,
    url: "/project-info",
    children: [],
    divider: false
  },
  {
    name: 'Ход строительства',
    icon:  <Tool size={ICON_SIZE} />,
    url: "/progress",
    children: [],
    divider: false,
  },
  {
    name: 'Планировки',
    icon:  <Layout size={ICON_SIZE} />,
    url: "/plans",
    children: [],
    divider: false
  },
  {
    name: 'Видео',
    icon:  <Youtube size={ICON_SIZE} />,
    url: "/video",
    children: [],
    divider: false,
  },
  {
    name: 'А так же вас ждет',
    icon:  <Grid size={ICON_SIZE} />,
    url: "/advantages",
    children: [],
    divider: false
  },
  {
    name: 'Документация',
    icon:  <BookOpen size={ICON_SIZE} />,
    url: "/docs",
    children: [],
    divider: false
  },
  {
    name: "Настройки",
    icon: null,
    url: null,
    children: null,
    divider: true,
  },
  {
    name: 'Подмены',
    icon:  <RefreshCw size={ICON_SIZE} />,
    url: "/titles",
    children: [],
    divider: false
  },
  {
    name: 'Параметры',
    icon:  <Settings size={ICON_SIZE} />,
    url: "/settings",
    children: [],
    divider: false
  },
  {
    name: "",
    icon: null,
    url: null,
    children: null,
    divider: true,
  },
  {
    name: "",
    icon: null,
    url: null,
    children: null,
    divider: true,
  },
  {
    name: 'Выход',
    icon:  <LogOut size={ICON_SIZE} />,
    url: "/logout",
    children: [],
    divider: false
  },
]

export default elements