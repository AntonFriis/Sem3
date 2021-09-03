using System;
using System.IO;
using Xunit;

namespace HelloWorld.Tests
{
    public class HelloWorldTest
    {
        [Fact]
        public void Main_prints_Hello_World()
        {
            // Arrange
            var writer = new StringWriter();
            
            Console.SetOut(writer);
            
            // Act
            HelloWorld.Main(new string[0]);
            
            // Assert
            var output = writer.GetStringBuilder().ToString().Trim();
            Assert.Equal("Hello World!", output);
        }
    }
}