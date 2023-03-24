import {
  Component,
  createEffect,
  createSignal,
  onCleanup,
  For,
} from "solid-js";
import styles from "./Slideshow.module.css";
import useSlideshow from "../../stores/useSlideshow";
import { Video, Image } from "../../types/models";

type Media = Image | Video;

const Slideshow: Component = () => {
  const { getMedia } = useSlideshow();
  const [currentSlide, setCurrentSlide] = createSignal(0);

  createEffect(() => {
    const currentMedia = getMedia()[currentSlide()];
    if (currentMedia.type === "image") {
      let currentImage = currentMedia as Image;
      const duration = currentImage.duration || 10;
      const interval = duration * 1000;
      const timer = setTimeout(() => {
        updateActiveImage(1);
      }, interval);
      onCleanup(() => clearTimeout(timer));
    } else if (currentMedia.type === "video") {
      const videoElement = document.getElementById(
        `slide-video-${currentMedia.id}`
      ) as HTMLVideoElement;
      videoElement.play();
      const onEnded = () => {
        videoElement.removeEventListener("ended", onEnded);
        updateActiveImage(1);
      };
      videoElement.addEventListener("ended", onEnded);
      onCleanup(() => {
        videoElement.removeEventListener("ended", onEnded);
        videoElement.pause();
        videoElement.currentTime = 0;
      });
    }
  });

  const updateActiveImage = (index: number) => {
    const nextSlide = currentSlide() + index;
    if (nextSlide < 0) {
      setCurrentSlide(getMedia().length - 1);
    } else if (nextSlide >= getMedia().length) {
      setCurrentSlide(0);
    } else {
      setCurrentSlide(nextSlide);
    }
  };

  return (
    <div class={styles.Slideshow}>
      <div class={styles.Slideshow__container}>
        <For each={getMedia()}>
          {(media: Media) => (
            <div
              class={styles.Slideshow__media}
              style={{
                display: currentSlide() === media.order ? "block" : "none",
              }}
            >
              {media.type === "image" ? (
                <div class={styles.Slideshow__image}>
                  <img src={media.url} alt={media.caption} />
                </div>
              ) : (
                <div class={styles.Slideshow__video}>
                  <video id={`slide-video-${media.id}`} src={media.url} muted autoplay />
                </div>
              )}
            </div>
          )}
        </For>
      </div>
    </div>
  );
};

export default Slideshow;
