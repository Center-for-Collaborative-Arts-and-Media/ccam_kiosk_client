import { SliderShowType } from "../types/models";
import { createStore } from "solid-js/store";

const useSlideshow = () => {
  const [slideshow, setSlideshow] = createStore<SliderShowType>({
    id: "asdkjhfyend",
    title: "ccam_slideshow",
    text_overlay: "",
    media: {
      id: "djfiiue",
      videos: [
        // {
        //   id: "djjfs783",
        //   type: "video",
        //   order: 1,
        //   caption: "",
        //   url: "https://mdn.github.io/dom-examples/canvas/chroma-keying/media/video.mp4",
        // },
        // {
        //   id: "djjfs784",
        //   type: "video",
        //   order: 3,
        //   caption: "",
        //   url: "https://samplelib.com/lib/preview/mp4/sample-30s.mp4",
        // },
      ],
      images: [
        {
          id: "dyenfs783",
          type: "image",
          order: 0,
          caption: "",
          url: "https://ccam.yale.edu/sites/default/files/styles/homepage_event_image/public/event-images/da482bf6-7a42-d9df-f619-cf8753fba624_4.png?itok=Kfp4wN_U",
          duration: 240,
        },
        {
          id: "djjfs784",
          type: "image",
          order: 2,
          caption: "",
          url: "https://ccam.yale.edu/sites/default/files/styles/homepage_event_image/public/event-images/ae3976fa-cd3e-c9c5-b1d4-0a11266503ef.png?itok=RqK0MrQS",
          duration: 240,
        },
      ],
    },
    archive: false,
  });

  const getMedia = () => {
    const img = Array.from(slideshow.media.images).sort(
      (a, b) => a.order - b.order
    );
    const vid = Array.from(slideshow.media.videos).sort(
      (a, b) => a.order - b.order
    );
    return [...img, ...vid].sort((a, b) => a.order - b.order);
  };

  return { slideshow, setSlideshow, getMedia };
};

export default useSlideshow;
