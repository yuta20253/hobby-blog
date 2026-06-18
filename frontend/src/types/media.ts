export type Media = {
  id: number;
  type: MediaType;
  file_path: string;
  file_name: string;
};

type MediaType = "image" | "video";
