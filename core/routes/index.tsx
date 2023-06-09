import LogoSection from "../components/Logo.tsx";
import { Head } from "$fresh/runtime.ts";

export default function App() {
  return (
    <>
      <Head>
        <meta charSet="UTF-8" />

        <meta content="width=device-width, initial-scale=1" name="viewport" />

        <link
          rel="icon"
          type="image/svg"
          href="https://cdn-botway.deno.dev/simple/icon.svg"
        />

        <title>📡 Get Latest Release</title>
      </Head>
      <main
        className="flex flex-col md:flex-row-reverse md:h-screen"
        style={{ background: "#13111c" }}
      >
        <LogoSection />

        <section className="justify-center px-4 md:px-0 md:flex md:w-2/3 md:border-r border-gray-800">
          <div className="w-full max-w-sm py-4 mx-auto my-auto min-w-min md:py-9 md:w-7/12">
            <h3 className="font-medium md:text-xl pt-3 text-white">
              Welcome, this is the <code>get-latest api</code>, please use{" "}
              <code>/:user/:repo/</code> route to get latest version of a repo
              📡
            </h3>
          </div>
        </section>
      </main>
    </>
  );
}
