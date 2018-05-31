using System;
using System.Collections.Generic;

namespace Mvc4.Razor2.Models
{
    public class AuthorModel
    {
        public int Id { get; set; }
        public DateTime BirthDate { get; set; }
        public DateTime? DateOfDeath { get; set; }
        public string FirstName { get; set; }
        public string LastName { get; set; }
        public string Tagline { get; set; }
        public string About { get; set; }

        public List<Book> Books { get; set; }
    }
}