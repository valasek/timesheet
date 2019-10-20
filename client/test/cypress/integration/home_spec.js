// DONE
describe('Timesheet home page contains all headers and some data', function() {
  it('Checks headers', function() {
    cy.visit('/')
    cy.get('div').contains('Top 10 projects in')
    cy.get('div').contains('Managed data')
  })

  it('Top 10 projects are shown', function () {
    cy.visit('/')
    cy.get('canvas').should('be.visible')
  })

  it('Top 10 projects table contains at least 5 rows', function () {
    cy.visit('/')
    cy.get('table').first().should('be.visible')
    cy.get('table').first().get('th').contains('Project')
    cy.get('table').first().get('th').contains('Hours')
    cy.get('tbody').children().should('have.length.greaterThan', 5)
  })

  it('Managed data table contains correct header and at least 1 record per each entity', function () {
    cy.visit('/')
    cy.get('table').last().should('be.visible')
    cy.get('table').last().get('th').contains('# of total records')
    cy.get('tbody').last().children().should('have.length', 4)
    // check reported records
    cy.get('tbody').last().children().eq(0).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
    // check projects
    cy.get('tbody').last().children().eq(1).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
    // check rates
    cy.get('tbody').last().children().eq(2).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
    // check consultants
    cy.get('tbody').last().children().eq(3).children().eq(1).invoke('text').then((text => {
      expect(text).to.not.eq('0')
    }))
  })
})
