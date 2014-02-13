using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Threading;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Mvc4.Razor2.Models;
using Mvc4.Razor2.Views.Authors;
using RazorGenerator.Testing;

namespace Mvc4.Razor2.Tests
{
    [TestClass]
    public class UnitTest1
    {
        private AuthorPage _view;
        private AuthorModel _author;

        [TestInitialize]
        public void WarmUp()
        {
            _view = new AuthorPage();
            _author = new AuthorModel { FirstName = "Alex", Books = new List<Book>() };
            for (int i = 1; i <= 1000; i++)
            {
                _author.Books.Add(new Book()
                {
                    Id = i,
                    Title = string.Format("Book #{0}", i),
                    FirstTimePublished = DateTime.Now.AddDays(i),
                    Price = 100 + (decimal)0.1 * i
                });
            }
            _view.Render(_author);
        }

        [TestMethod]
        public void TestMethod0_WarmUp()
        {
        }

        [TestMethod]
        public void TestMethod1_Render()
        {
            Process.GetCurrentProcess().ProcessorAffinity = new IntPtr(1);
            Process.GetCurrentProcess().PriorityClass = ProcessPriorityClass.High;
            Thread.CurrentThread.Priority = ThreadPriority.Highest;
            const int rounds = 1009; //2003;
            string content = string.Empty;

            var min = new Stopwatch();
            var max = new Stopwatch();
            int maxRound = -1, minRound = -1;
            max.Start();
            max.Stop();
            var i = 0;
            //warmup
            min.Start();
            while(i++<100)
                _view.Render(_author);
            min.Stop();

            i = 0;
            var stopwatchTotal = Stopwatch.StartNew();
            //var started = DateTime.Now;
            while(i++<rounds)
            {
                var stopwatchSingle = Stopwatch.StartNew();
                content = _view.Render(_author);
                stopwatchSingle.Stop();
                if (min.ElapsedTicks == 0 || stopwatchSingle.ElapsedTicks < min.ElapsedTicks)
                {
                    min = stopwatchSingle;
                    minRound = i;
                }
                if (stopwatchSingle.ElapsedTicks > max.ElapsedTicks)
                {
                    max = stopwatchSingle;
                    maxRound = i;
                }
            }
            stopwatchTotal.Stop();

            Console.WriteLine("Average: {0} microseconds, Min: {1} microseconds, MinRound: {2}, Max: {3} microseconds, MaxRound: {4}, TotalRounds: {5}",
                // ReSharper disable once PossibleLossOfFraction
                (stopwatchTotal.ElapsedTicks / rounds / 10),
                // ReSharper disable once PossibleLossOfFraction
                Math.Round((decimal) (min.ElapsedTicks / 10), 3),
                minRound,
                // ReSharper disable once PossibleLossOfFraction
                Math.Round((decimal)(max.ElapsedTicks / 10), 3),
                maxRound,
                rounds
                );
            Console.WriteLine(content);
        }
    }
}
