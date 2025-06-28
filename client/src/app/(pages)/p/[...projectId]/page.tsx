import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import React from "react";
import SceneDescriptionCard from "./_components/sceneDescriptionCard";
import { data } from "@/lib/utils";
import { Accordion } from "@/components/ui/accordion";
import { Button } from "@/components/ui/button";

function Project() {

    
  return (
    <div className=" flex h-full  items-center justify-center ">
      <div className=" flex flex-col gap-4  h-full  w-[50%] p-2">
        <div>
          <h1 className=" text-2xl font-semibold">Edit Plan</h1>
        </div>

        <div>
          <h2>Project Metadata</h2>
          <Separator />

          <div className=" flex flex-col  gap-1 mt-1">
            <div className=" flex gap-2 flex-col">
              <label className="  text-xs">Title</label>
              <Input value={data.metaData.title} />
            </div>
            <div className=" flex flex-col gap-2">
              <label className=" text-xs">Description</label>
              <Input value={data.metaData.description} />
            </div>
          </div>
        </div>

        <div>
          <h2>Scenes Description</h2>
          <Separator />
          <Accordion type="single" collapsible>
            {data.scenesData.map((scene) => (
              <>
                <SceneDescriptionCard
                  title={scene.sceneTitle}
                  description={scene.instruction}
                  sequence={scene.sequence}
                  key={scene.sequence}
                />
                {/* <Separator /> */}
              </>
            ))}
          </Accordion>
        </div>
        <Button>Proceed to generate</Button>
      </div>
    </div>
  );
}

export default Project;
