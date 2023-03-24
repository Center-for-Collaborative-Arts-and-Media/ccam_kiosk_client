export type SliderShowType = {
  id: string;
  title: string;
  text_overlay: string;
  media: Media;
  archive: boolean;
};

type Video = {
  id: string;
  type: string;
  order: number;
  caption: string;
  url: string;
};

type Image = {
  id: string;
  type: string;
  order: number;
  caption: string;
  url: string;
  duration: number;
};

type Media = {
  id: string;
  videos: Video[];
  images: Image[];
};
