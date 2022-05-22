import { Layers, Tool, Layout, Youtube, Grid, BookOpen, Server } from 'react-feather';

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
]

export default elements