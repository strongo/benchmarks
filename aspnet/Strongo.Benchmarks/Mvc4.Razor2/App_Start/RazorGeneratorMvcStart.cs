using System.Web;
using System.Web.Mvc;
using System.Web.WebPages;
using RazorGenerator.Mvc;

[assembly: WebActivatorEx.PostApplicationStartMethod(typeof(Mvc4.Razor2.App_Start.RazorGeneratorMvcStart), "Start")]

namespace Mvc4.Razor2.App_Start {
    public static class RazorGeneratorMvcStart {
        public static void Start() {
            var engine = new PrecompiledMvcEngine(typeof(RazorGeneratorMvcStart).Assembly) {
                UsePhysicalViewsIfNewer = HttpContext.Current.Request.IsLocal
            };

            engine.PreemptPhysicalFiles = true;

            ViewEngines.Engines.Add(engine);

            // StartPage lookups are done by WebPages. 
            VirtualPathFactoryManager.RegisterVirtualPathFactory(engine);
        }
    }
}
