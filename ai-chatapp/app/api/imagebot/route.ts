import { NextResponse } from "next/server";
import { Configuration, OpenAIApi } from "openai";

const configuration = new Configuration({
  apiKey: process.env.OPENAI_API_KEY,
});

const openai = new OpenAIApi(configuration);
export async function POST(req: Request) {
  try {
    const body = await req.json();

    const { messages } = body;

    if (!configuration.apiKey) {
      return new NextResponse("API Key required", { status: 401 });
    }

    if (!messages) {
      return new NextResponse("Msg required", { status: 400 });
    }

    const response = await openai.createImage({
      prompt: messages,
      n: 2,
      size: "512x512",
    });

    return NextResponse.json(response.data.data);
  } catch (error) {
    console.log(error);
    return new NextResponse("An error occurred", { status: 500 });
  }
}
