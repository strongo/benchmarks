using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.Mvc;

namespace Mvc4.Razor2.Controllers
{
    using Models;

    public class AuthorsController : Controller
    {
        //
        // GET: /Authors/
        //public ActionResult Index()
        //{
        //    return View();
        //}

        //
        // GET: /Authors/AuthorPage/5

        private static readonly AuthorModel TestAuthor;

        static AuthorsController()
        {
            TestAuthor = new AuthorModel
            {
                Id = 1,
                FirstName = "John",
                LastName = "Smith",
                Tagline = "Test author.",
                About = "This is just a test author who never lived.",
                Books = new List<Book>()
            };

            for (int i = 0; i < 100; i++)
            {
                TestAuthor.Books.Add(new Book()
                {
                    Id = i,
                    Title = string.Format("Book #{0}", i),
                    FirstTimePublished = DateTime.Now.AddDays(i),
                    Price = 100 + (decimal)0.1*i
                });
            }
        }

        public ActionResult AuthorPage(int id)
        {
            return View(TestAuthor);
        }
    }
}
